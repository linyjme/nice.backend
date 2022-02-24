package tag_handler

import (
	"niceBackend/internal/pkg/core"
)

type loginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  dataResponse  `json:"result"`
}
type dataResponse struct {
	Data  []int  `json:"data"`
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
		res := new(loginResponse)
		res.Status = "success"
		res.Message = "登陆成功"
		res.Result = dataResponse{
			Data: []int{},
		}
		c.Payload(res)
	}
}
