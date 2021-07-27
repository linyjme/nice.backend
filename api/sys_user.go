package api

import (
	"asyncClient/common/response"
	"github.com/gin-gonic/gin"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	//var l request.Login
	//_ = c.ShouldBindJSON(&l)
	//if err := utils.Verify(l, utils.LoginVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//if store.Verify(l.CaptchaId, l.Captcha, true) {
	//	u := &model.SysUser{Username: l.Username, Password: l.Password}
	//	if err, user := service.Login(u); err != nil {
	//		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
	//		response.FailWithMessage("用户名不存在或者密码错误", c)
	//	} else {
	//		tokenNext(c, *user)
	//	}
	//} else {
	//	response.FailWithMessage("验证码错误", c)
	//}
	response.FailWithMessage("验证码错误", c)
}
