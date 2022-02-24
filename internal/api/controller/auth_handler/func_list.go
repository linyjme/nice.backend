package auth_handler

import "niceBackend/internal/pkg/core"

type authRequest struct {
	Password string `json:"password"` // 用户 账号密码
	Account  string `json:"account"`  // 用户 账号密码
}

type authResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  Token  `json:"result"`
}

// List 管理员列表
// @Summary 管理员列表
// @Description 管理员列表
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param page query int true "第几页" default(1)
// @Param page_size query int true "每页显示条数" default(10)
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param mobile query string false "手机号"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [get]
// @Security LoginToken
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		res := new(authResponse)
		res.Status = "success"
		res.Message = "登陆成功"
		res.Result = Token{
			AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXIiOiJyb290In0sImlhdCI6MTY0NTY5NTkxMiwiZXhwIjoxNjQ1Njk5NTEyfQ.o4xs9FrYHDam5YhK2hmDsPO0qLhsajX3mCnhPwH0wyY",
			ExpiresIn:   3600,
		}
		c.Payload(res)
	}
}
