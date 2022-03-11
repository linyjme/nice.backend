package admin

import (
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/repository/db_repo/admin_repo"
)

type SearchOneData struct {
	Id       int32  // 用户ID
	Account string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
	IsUsed   int32  // 是否启用 1:是  -1:否
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *admin_repo.Admin, err error) {
	return info, nil

}
