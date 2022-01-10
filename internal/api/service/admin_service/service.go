package admin_service

var _ Service = (*service)(nil)

type Service interface {
	i()
}

type service struct {
}

func New() Service {
	return &service{
	}
}

func (s *service) i() {}
