package middleware

import (
	"niceBackend/internal/api/service/admin"
	"niceBackend/internal/pkg/cache"
	"niceBackend/internal/pkg/core"
	"niceBackend/internal/pkg/db"
	"niceBackend/pkg/errno"

	"go.uber.org/zap"
)

var _ Middleware = (*middleware)(nil)

type Middleware interface {
	// i 为了避免被其他包实现
	i()

	// Jwt 中间件
	Jwt(ctx core.Context) (userId int64, userName string, err errno.Error)

	// Resubmit 中间件
	Resubmit() core.HandlerFunc

	// DisableLog 不记录日志
	DisableLog() core.HandlerFunc
}

type middleware struct {
	logger       *zap.Logger
	cache        cache.Repo
	db           db.Repo
	adminService admin.Service
}

func New(logger *zap.Logger, cache cache.Repo, db db.Repo) Middleware {
	return &middleware{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func (m *middleware) i() {}
