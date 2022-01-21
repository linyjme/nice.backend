package comment_handler

import (
	"niceBackend/internal/api/service/user_service"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
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
