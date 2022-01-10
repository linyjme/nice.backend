package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niceBackend/common/global"
	"niceBackend/common/transform/request"
	"niceBackend/common/transform/response"
	"niceBackend/internal/api/repository/db_repo/user_repo"
	"niceBackend/internal/api/service/user_service"
	"niceBackend/pkg"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Register(c *gin.Context) {
	var r request.Register
	_ = c.ShouldBindJSON(&r)
	if err := pkg.Verify(r, pkg.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &user_repo.User{Account: r.Account, Password: r.Password}
	err, userReturn := user_service.Register(*user)
	if err != nil {
		global.NiceLog.Error("注册失败!", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}
