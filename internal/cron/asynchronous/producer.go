package asynchronous

import (
	"container/list"

	"niceBackend/common/global"
)

func DistributeTask(task list.List) {
	for e := task.Front(); e != nil; e = e.Next() {
		global.AsyncChan.PutChan(e.Value.(map[string]interface{}))
	}

}
