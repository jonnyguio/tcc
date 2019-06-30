package main

import (
	"fmt"
	"sync"
)

type Reader struct {
	ID      int
	Readers *sync.Mutex
	Writers *sync.Mutex
	Count   chan int
}

func (r *Reader) Start(stream []byte, size int) {
	for {
		r.Readers.Lock()
		count := <-r.Count
		if count == 0 {
			r.Writers.Lock()
		}
		r.Count <- count + 1
		r.Readers.Unlock()

		fmt.Printf("Reader %d read byte string %v\n", r.ID, stream)
		count = <-r.Count
		if count == 1 {
			r.Writers.Unlock()
		}
		r.Count <- count - 1
	}
}
