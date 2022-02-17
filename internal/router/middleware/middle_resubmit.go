package middleware

import (
	"net/http"
	"niceBackend/config"
	"time"

	"niceBackend/configs"
	"niceBackend/internal/pkg/cache"
	"niceBackend/internal/pkg/code"
	"niceBackend/internal/pkg/core"
	"niceBackend/pkg/errno"
	"niceBackend/pkg/errors"
	"niceBackend/pkg/token"
)

const reSubmitMark = "1"

func (m *middleware) Resubmit() core.HandlerFunc {
	return func(c core.Context) {
		cfg := configs.Get().URLToken

		tokenString, err := token.New(cfg.Secret).UrlSign(c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.UrlSignError,
				code.Text(code.UrlSignError)).WithErr(err),
			)
			return
		}

		redisKey := config.RedisKeyPrefixRequestID + tokenString
		if !m.cache.Exists(redisKey) {
			err = m.cache.Set(redisKey, reSubmitMark, time.Minute*cfg.ExpireDuration)
			if err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.CacheSetError,
					code.Text(code.CacheSetError)).WithErr(err),
				)
				return
			}

			return
		}

		redisValue, err := m.cache.Get(redisKey, cache.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithErr(err),
			)
			return
		}

		if redisValue == reSubmitMark {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ResubmitError,
				code.Text(code.ResubmitError)).WithErr(errors.New("resubmit")),
			)
			return
		}

		return
	}
}
