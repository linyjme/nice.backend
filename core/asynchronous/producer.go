package asynchronous

import (
	"container/list"
)

func DistributeTask(task list.List) {

}

func Producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
