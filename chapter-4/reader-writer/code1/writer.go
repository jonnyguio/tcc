package main

import (
	"crypto/rand"
	"fmt"
	"sync"
)

type Writer struct {
	ID    int
	Mutex *sync.RWMutex
}

func (w *Writer) start(stream []byte, size int) {
	for {
		w.Mutex.Lock()
		fmt.Printf("Writer %d wrinting byte string %v\n", w.ID, stream)
		rand.Read(stream)
		// time.Sleep(50 * time.Millisecond)
		w.Mutex.Unlock()
	}
}
