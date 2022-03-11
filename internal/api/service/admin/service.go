package admin

import (
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/repository/db_repo/admin_repo"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Detail(ctx core.Context, searchOneData *SearchOneData)  (info *admin_repo.Admin, err error)
}

type service struct {
}

func New() Service {
	return &service{
	}
}

func (s *service) i() {}
