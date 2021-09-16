package asynchronous

import (
	"container/list"
	"sync"
	"sync/atomic"
)

// 等待协程的数量
var waitRoutine int32 = 0

// Itemer ..
type Itemer interface {
	Process()
}

// TaskQueue ..
type TaskQueue struct {
	tl *list.List

	size    int // 队列大小
	routine int // 开启协程的数量

	sync.Mutex
	*sync.Cond
}

// NewTaskQueue ..
func NewTaskQueue(size int, routine int) *TaskQueue {
	if size <= 0 {
		size = 4
	}
	if routine <= 0 {
		routine = 4
	}

	taskQueue := &TaskQueue{
		size:    size,
		routine: routine,
	}
	taskQueue.Cond = sync.NewCond(&taskQueue.Mutex)
	taskQueue.tl = &list.List{}
	return taskQueue
}

// Run ..
func (taskQueue *TaskQueue) Run() {
	for index := 0; index < taskQueue.routine; index++ {
		go process(taskQueue)
	}
}

// PushItem ..
func (taskQueue *TaskQueue) PushItem(item Itemer) bool {

	taskQueue.Mutex.Lock()
	if taskQueue.tl.Len() >= taskQueue.size {
		taskQueue.Mutex.Unlock()
		return false
	} else {
		taskQueue.tl.PushBack(item)
	}
	taskQueue.Mutex.Unlock()

	if atomic.LoadInt32(&waitRoutine) > 0 {
		taskQueue.Cond.Signal()
	}

	return true
}

func process(taskQueue *TaskQueue) {
	for {
		taskQueue.Mutex.Lock()
		for taskQueue.tl.Len() == 0 {

			atomic.AddInt32(&waitRoutine, 1)
			taskQueue.Cond.Wait()
			atomic.AddInt32(&waitRoutine, -1)
		}

		item := taskQueue.tl.Front()
		taskQueue.tl.Remove(item)
		taskQueue.Mutex.Unlock()

		item.Value.(Itemer).Process()
	}
}
