package middleware

import "niceBackend/internal/pkg/core"

func (m *middleware) DisableLog() core.HandlerFunc {
	return func(c core.Context) {
		core.DisableTrace(c)
	}
}
