package expansion_handler

import (
	"niceBackend/internal/api/service/user_service"
	"niceBackend/internal/pkg/core"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 管理员登录
	// @Tags API.admin
	// @Router /api/tag [get]
	List() core.HandlerFunc
	Login() core.HandlerFunc
}

type handler struct {
	userService user_service.Service
}

func New() Handler {
	return &handler{
		userService: user_service.New(),
	}
}

func (h *handler) i() {}
