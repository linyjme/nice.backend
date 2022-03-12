package initialize

import (
	"niceBackend/common/transform"
)

func InitAsync() *transform.AsyncChan {
	return &transform.AsyncChan{
		Product:   make(chan map[string]interface{}, 1000),
		CloseChan: make(chan bool),
		IsClosed:  false,
	}
}
