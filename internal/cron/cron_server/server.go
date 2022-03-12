package cron_server

import (
	"sync"

	"niceBackend/internal/pkg/cache"
	"niceBackend/pkg/errors"

	"github.com/jakecoffman/cron"
	"go.uber.org/zap"
)

var _ Server = (*server)(nil)

type taskCount struct {
	wg   sync.WaitGroup
	exit chan struct{}
}

func (tc *taskCount) i() {}

func (tc *taskCount) Add() {
	tc.wg.Add(1)
}

func (tc *taskCount) Done() {
	tc.wg.Done()
}

func (tc *taskCount) Exit() {
	tc.wg.Done()
	<-tc.exit
}

func (tc *taskCount) Wait() {
	tc.Add()
	tc.wg.Wait()
	close(tc.exit)
}

type server struct {
	logger    *zap.Logger
	cache     cache.Repo
	cron      *cron.Cron
	taskCount *taskCount
}

type Server interface {
	i()

	// Start 启动 cron 服务
	Start()
}

func New(logger *zap.Logger, cache cache.Repo) (Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	if cache == nil {
		return nil, errors.New("cache required")
	}

	return &server{
		logger: logger,
		cache:  cache,
		cron:   cron.New(),
		taskCount: &taskCount{
			wg:   sync.WaitGroup{},
			exit: make(chan struct{}),
		},
	}, nil
}

func (s *server) i() {}
