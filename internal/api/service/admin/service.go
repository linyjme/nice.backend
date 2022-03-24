package admin

import (
	"niceBackend/common/global"
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/repository/mysql/model"
	"niceBackend/internal/repository/mysql/query"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	FindByAccountAndPassword(ctx core.Context, account string, password string) (result *model.Administrator, err error)
}

type service struct {
	qu *query.Query
}

func New() Service {
	return &service{
		qu: query.Use(global.NiceDb),
	}
}

func (s *service) i() {}
