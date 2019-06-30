package main

import (
	"fmt"
	"sync"
)

type Reader struct {
	ID      int
	Readers chan int
	Writers *sync.Mutex
}

func (r *Reader) start(stream []byte, size int) {
	for {
		count := <-r.Readers
		if count == 0 {
			fmt.Printf("Reader %d getting lock\n", r.ID)
			r.Writers.Lock()
			fmt.Printf("Reader %d got lock\n", r.ID)
		}
		r.Readers <- count + 1

		fmt.Printf("Reader %d read byte string %v\n", r.ID, stream)
		count = <-r.Readers
		if count == 1 {
			r.Writers.Unlock()
		}
		r.Readers <- count - 1
	}
}
