package middleware

import (
	"net/http"
	"niceBackend/configs"
	"niceBackend/internal/pkg/code"
	"niceBackend/internal/pkg/core"
	"niceBackend/pkg/errno"
	"niceBackend/pkg/errors"
	"niceBackend/pkg/token"
)

func (m *middleware) Jwt(ctx core.Context) (userId int64, userName string, err errno.Error) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中缺少 Authorization 参数"))

		return
	}

	cfg := configs.Get().JWT
	claims, errParse := token.New(cfg.Secret).JwtParse(auth)
	if errParse != nil {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errParse)

		return
	}

	userId = claims.UserID
	if userId <= 0 {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("claims.UserID <= 0 "))

		return
	}
	userName = claims.UserName
	return
}
