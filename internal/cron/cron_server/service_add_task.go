package cron_server

import (
	"strings"

	"niceBackend/internal/repository/db_repo/cron_task_repo"

	"github.com/spf13/cast"
)

func (s *server) AddTask(task *cron_task_repo.CronTask) {
	spec := "0 " + strings.TrimSpace(task.Spec)
	name := cast.ToString(task.Id)

	s.cron.AddFunc(spec, s.AddJob(task), name)
}
