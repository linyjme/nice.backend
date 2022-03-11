package statistic

import "niceBackend/internal/pkg/core"

var _ Service = (*service)(nil)

type Service interface {
	i()
	Create(ctx core.Context) (id int32, err error)
}

type service struct {
}

func New() Service {
	return &service{
	}
}

func (s *service) i() {}
