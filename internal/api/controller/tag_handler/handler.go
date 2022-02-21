package tag_handler

import (
	"niceBackend/internal/api/service/tag_service"
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
	tagService tag_service.Service
}

func New() Handler {
	return &handler{
		tagService: tag_service.New(),
	}
}

func (h *handler) i() {}
