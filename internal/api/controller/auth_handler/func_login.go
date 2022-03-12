package auth_handler

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"sync"
	"time"

	"niceBackend/common/global"
	"niceBackend/common/transform/request"
	"niceBackend/common/transform/response"
	"niceBackend/config"
	"niceBackend/internal/api/service/admin"
	"niceBackend/internal/middleware"
	"niceBackend/internal/pkg/code"
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/pkg/password"
	"niceBackend/internal/repository/db_repo/admin_repo"
	"niceBackend/pkg"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
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
		if err := pkg.Verify(req, pkg.LoginVerify); err != nil {
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
		pass := pkg.DecodeBase64(req.Password)
		account := req.Account
		if account == "" {
			account = pass
		}
		var shaKey string
		if userLoginCache != nil {
			h := sha256.New()
			h.Write(pkg.StringToBytes(account + "|" + pass))
			shaKey = string(h.Sum(nil))
			u, ok := userLoginCache.Get(shaKey)
			if ok {
				res.Result = u.(Token)
				c.Payload(res)
				return
			}
		}
		searchOneData := new(admin.SearchOneData)
		searchOneData.Account = account
		searchOneData.Password = password.GeneratePassword(pass)
		info, _ := h.adminService.Detail(c, searchOneData)
		fmt.Println(info)
		//h.userService
		//u := &model.User{Account: account, Password: password}
		//if err, user := service.Login(u); err != nil {
		//	global.NiceLog.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		//	response.FailWithCode(4002, c)
		//} else {
		//	// 颁发token
		//	tokenNext(c, *user)
		//}
		res.Status = "success"
		res.Message = "登陆成功"
		token := Token{
			AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXIiOiJyb290In0sImlhdCI6MTY0NTY5NTkxMiwiZXhwIjoxNjQ1Njk5NTEyfQ.o4xs9FrYHDam5YhK2hmDsPO0qLhsajX3mCnhPwH0wyY",
			ExpiresIn:   3600,
		}
		res.Result = token
		c.Payload(res)
		if userLoginCache != nil {
			userLoginCache.SetDefault(shaKey, token)
		}
	}

}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user admin_repo.Admin) {
	j := &middleware.JWT{SigningKey: []byte(global.NiceConfig.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		Account:    user.Account,
		BufferTime: global.NiceConfig.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.NiceConfig.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                              // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.NiceLog.Error("获取token失败!", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(response.LoginResponse{
		Admin:     user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
	return
}
