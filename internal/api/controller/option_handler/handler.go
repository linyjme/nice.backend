package option_handler

import (
	"niceBackend/internal/pkg/core"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	// List 管理员登录
	// @Tags API.admin
	// @Router /api/tag [get]
	List() core.HandlerFunc
}

type handler struct {
}

func New() Handler {
	return &handler{}
}

func (h *handler) i() {}
