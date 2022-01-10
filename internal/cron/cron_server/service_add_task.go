package cron_server

import (
	"github.com/spf13/cast"
	"niceBackend/internal/api/repository/db_repo/cron_task_repo"
	"strings"
)

func (s *server) AddTask(task *cron_task_repo.CronTask) {
	spec := "0 " + strings.TrimSpace(task.Spec)
	name := cast.ToString(task.Id)

	s.cron.AddFunc(spec, s.AddJob(task), name)
}
