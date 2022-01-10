package asynchronous

import "niceBackend/common/transform"

func StopChan(p *transform.AsyncChan) {
	p.Close()
}
