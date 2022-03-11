package user_handler

import (
	"niceBackend/internal/api/service/admin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
}

type handler struct {
	userService admin.Service
}

func New() Handler {
	return &handler{
		userService: admin.New(),
	}
}

func (h *handler) i() {}
