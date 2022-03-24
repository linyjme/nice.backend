package asynchronous

import (
	"niceBackend/internal/pkg/async"
)

func StopChan(p *async.AsyncChan) {
	p.Close()
}
