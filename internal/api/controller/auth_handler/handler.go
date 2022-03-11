package auth_handler

import (
	"niceBackend/internal/api/service/admin"
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
	adminService admin.Service
}

func New() Handler {
	return &handler{
		adminService: admin.New(),
	}
}

func (h *handler) i() {}
