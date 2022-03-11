package tag_handler

import (
	"niceBackend/internal/api/service/tag"
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
	tagService tag.Service
}

func New() Handler {
	return &handler{
		tagService: tag.New(),
	}
}

func (h *handler) i() {}
