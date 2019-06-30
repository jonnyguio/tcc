package reader

import (
	"fmt"
	"sync"
)

type Reader struct {
	ID      int
	Readers *sync.Mutex
	Writers *sync.Mutex
}

var count = 0
var local = &sync.Mutex{}

func (r *Reader) Start(stream []byte, size int) {
	for {
		r.Readers.Lock()
		local.Lock()
		count++
		if count == 1 {
			r.Writers.Lock()
		}
		local.Unlock()
		r.Readers.Unlock()

		fmt.Printf("Reader %d read byte string %v\n", r.ID, stream)
		local.Lock()
		count--
		if count == 0 {
			r.Writers.Unlock()
		}
		local.Unlock()
	}
}
