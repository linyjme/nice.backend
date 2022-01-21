package middleware

import (
	"niceBackend/internal/api/service/admin_service"
	"niceBackend/internal/api/service/authorized_service"
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

	// Signature 签名验证，对用签名算法 pkg/signature
	Signature() core.HandlerFunc

	// Token 签名验证，对登录用户的验证
	Token(ctx core.Context) (userId int64, userName string, err errno.Error)

	// RBAC 权限验证
	RBAC() core.HandlerFunc
}

type middleware struct {
	logger            *zap.Logger
	cache             cache.Repo
	db                db.Repo
	authorizedService authorized_service.Service
	adminService      admin_service.Service
}

func New(logger *zap.Logger, cache cache.Repo, db db.Repo) Middleware {
	return &middleware{
		logger:            logger,
		cache:             cache,
		db:                db,
		authorizedService: authorized_service.New(db, cache),
		adminService:      admin_service.New(db, cache),
	}
}

func (m *middleware) i() {}
