package async

import (
	"fmt"
	"sync"
)

func Init() *AsyncChan {
	return &AsyncChan{
		Product:   make(chan map[string]interface{}, 1000),
		CloseChan: make(chan bool),
		IsClosed:  false,
	}
}

type AsyncChan struct {
	Mutex     sync.Mutex
	Product   chan map[string]interface{}
	CloseChan chan bool
	IsClosed  bool
}

func (c *AsyncChan) PutChan(mapT map[string]interface{}) {
	c.Product <- mapT
}

//func (p *AsyncChan) Consumer() {
//	var t int
//	for true {
//		select {
//		case t = <-p.Product:
//			fmt.Println("Get:", t)
//			if t == 0 {
//				time.Sleep(1 * time.Second)
//			}
//		case p.IsClosed = <-p.CloseChan:
//			fmt.Println("finished")
//			close(p.CloseChan)
//			break
//		}
//		if p.IsClosed {
//			break
//		}
//	}
//}

func (p *AsyncChan) Close() {
	p.Mutex.Lock()
	if !p.IsClosed {
		fmt.Println("close chan")
		p.IsClosed = true
	}
	p.Mutex.Unlock()

}
