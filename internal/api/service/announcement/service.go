package announcement

import "niceBackend/internal/pkg/core"

var _ Service = (*service)(nil)

type Service interface {
	i()
	Create(ctx core.Context) (err error)
}

type service struct {
}

func New() Service {
	return &service{
	}
}

func (s *service) i() {}
