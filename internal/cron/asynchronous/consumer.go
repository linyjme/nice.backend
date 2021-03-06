package asynchronous

import (
	"fmt"
	"niceBackend/internal/pkg/async"

	"niceBackend/internal/cron/asynchronous/tasks"
)

func Consumer(p *async.AsyncChan) {
	var tMap map[string]interface{}
	for true {
		select {
		case tMap = <-p.Product:
			for k, v := range tMap {
				if k == "report_event" {
					fmt.Println(v)
					go tasks.ReportEventJob()
				}
			}
		case p.IsClosed = <-p.CloseChan:
			fmt.Println("close the chan ...")
			close(p.CloseChan)
			break
		default:
		}
		if p.IsClosed {
			break
		}

	}

}
