package interceptor

import (
	"go.uber.org/zap"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {

	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) Interceptor {
	return &interceptor{
		logger: logger,
	}
}

func (i *interceptor) i() {}
