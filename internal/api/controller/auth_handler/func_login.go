package auth_handler

import (
	"crypto/sha256"
	"net/http"
	"niceBackend/internal/repository/mysql/model"
	"niceBackend/pkg/util"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/patrickmn/go-cache"
	"niceBackend/common/transform/request"
	"niceBackend/config"
	"niceBackend/internal/middleware"
	"niceBackend/internal/pkg/code"
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/pkg/password"
)

// Admin login structure
type loginRequest struct {
	Password string `json:"password"` // 用户 账号密码
	Account  string `json:"account"`  // 用户 账号密码
}

type loginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  Token  `json:"result"`
}
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint16 `json:"expires_in"`
}

var (
	loginOnce      = sync.Once{}
	userLoginCache *cache.Cache
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		var req loginRequest
		res := new(loginResponse)
		_ = c.ShouldBindJSON(&req)
		if err := util.Verify(req, util.LoginVerify); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		loginOnce.Do(func() {
			if config.LoginCacheTime > 0 {
				userLoginCache = cache.New(time.Duration(config.LoginCacheTime)*time.Second, 2*time.Minute)
			}
		})
		pass := util.DecodeBase64(req.Password)
		account := req.Account
		if account == "" {
			account = pass
		}
		var shaKey string
		if userLoginCache != nil {
			h := sha256.New()
			h.Write(util.StringToBytes(account + "|" + pass))
			shaKey = string(h.Sum(nil))
			u, ok := userLoginCache.Get(shaKey)
			if ok {
				res.Result = u.(Token)
				c.Payload(res)
				return
			}
		}

		encryptionPassword := password.GeneratePassword(pass)
		info, err := h.adminService.FindByAccountAndPassword(c, account, encryptionPassword)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithError(err),
			)
			return
		}
		token := tokenNext(c, info)
		// 将用户信息记录到 Redis 中
		// 将可访问接口信息记录到 Redis 中
		//}
		res.Result = Token{
			AccessToken: token,
			ExpiresIn:   3600,
		}
		c.Payload(res)
		if userLoginCache != nil {
			userLoginCache.SetDefault(shaKey, token)
		}
	}

}

// 登录以后签发jwt
func tokenNext(c core.Context, user *model.Administrator) string {
	j := &middleware.JWT{SigningKey: []byte(config.GetConf().JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:       user.UID,
		ID:         user.ID,
		Account:    user.Account,
		BufferTime: config.GetConf().JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                             // 签名生效时间
			ExpiresAt: time.Now().Unix() + config.GetConf().JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                             // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.AbortWithError(core.Error(
			http.StatusBadRequest,
			code.AdminLoginError,
			code.Text(code.AdminLoginError)).WithError(err),
		)
		return token
	}
	//if err != nil {
	//	global.NiceLog.Error("获取token失败!", zap.Any("err", err))
	//	response.FailWithMessage("获取token失败", c)
	//	return ""
	//}
	return token
}
