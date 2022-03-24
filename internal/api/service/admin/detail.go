package admin

import (
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/repository/mysql/model"
)

func (s *service) FindByAccountAndPassword(ctx core.Context, account string, password string) (result *model.Administrator, err error) {
	ud := s.qu.Administrator.WithContext(ctx)
	result, err = ud.SimpleFindByAccountAndPassword(account, password)
	return result, err

}
