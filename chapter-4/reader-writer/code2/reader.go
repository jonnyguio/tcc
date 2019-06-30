package main

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

func (r *Reader) start(stream []byte, size int) {
	for {
		r.Readers.Lock()
		count++
		if count == 1 {
			r.Writers.Lock()
		}
		r.Readers.Unlock()

		fmt.Printf("Reader %d read byte string %v\n", r.ID, stream)
		r.Readers.Lock()
		count--
		if count == 0 {
			r.Writers.Unlock()
		}
		r.Readers.Unlock()
	}
}
