package asy_task

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"gm_server/models"
)

func RestoreRunnerStart(cardId int, data []models.TblCard) {
	//开启多核心
	runtime.GOMAXPROCS(runtime.NumCPU())

	//创建runner对象，设置超时时间
	runner := NewRunner(18 * time.Second)

	//total := len(data)/3 + 1
	//start := 0
	//end := 3
	//var tasks []func(id int) error
	//for i := 0; i < total; i++ {
	// if end > len(data) {
	//    end = len(data)
	// }
	// tasks = append(tasks,createRestoreTask(cardId, data[start:end]))
	// fmt.Println(tasks)
	// //添加运行的任务
	// start += 3
	// end += 3
	//}
	//
	//runner.Add(
	// tasks,
	//)

	runner.Add(
		createRestoreTask(cardId, data),
		//createTask(),
		//createTask(),
		//createTask(),
	)

	fmt.Println("异步执行任务")

	//开始执行任务
	if err := runner.Start(); err != nil {
		switch err {
		case ErrTimeout:
			fmt.Println("执行超时")
			os.Exit(1)
		case ErrInterrupt:
			fmt.Println("任务被中断")
			os.Exit(2)
		}
	}

	fmt.Println("执行结束")
}

//创建要执行的任务
func createRestoreTask(cardId int, data []models.TblCard) func(id int) error {
	return func(id int) error {
		fmt.Printf("正在执行%v个任务\n", id+1)
		fmt.Printf("一共%v条数据\n", len(data))
		fmt.Println(data)
		//模拟任务执行,sleep
		//time.Sleep(1 * time.Second)
		return nil
	}
}
