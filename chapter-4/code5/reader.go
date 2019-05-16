package main

import (
	"fmt"
	"sync"
)

type Reader struct {
	ID    int
	Mutex *sync.RWMutex
}

func (r *Reader) start(stream []byte, size int) {
	for {
		r.Mutex.RLock()
		fmt.Printf("Reader %d read byte string %v\n", r.ID, stream)
		// time.Sleep(time.Millisecond * 100)
		r.Mutex.RUnlock()
	}
}
