package cron_server

import (
	"fmt"
	"niceBackend/internal/api/repository/db_repo/cron_task_repo"
)

func (s *server) Start() {
	s.cron.Start()
	go s.taskCount.Wait()
	fmt.Println("start job")
	taskNum := 0
	listData := [...]*cron_task_repo.CronTask{

	}
	for _, item := range listData {
		s.AddTask(item)
		taskNum++
	}
}
