package main

import (
	"crypto/rand"
	"fmt"
	"sync"
)

type Writer struct {
	ID      int
	Writers *sync.Mutex
	Readers *sync.Mutex
	Count   chan int
}

func (w *Writer) Start(stream []byte, size int) {
	for {
		count := <-w.Count
		if count == 0 {
			w.Readers.Lock()
		}
		w.Count <- count + 1

		w.Writers.Lock()
		fmt.Printf("Writer %d writing byte string %v\n", w.ID, stream)
		rand.Read(stream)
		w.Writers.Unlock()

		count = <-w.Count
		if count == 1 {
			w.Readers.Unlock()
		}
		w.Count <- count - 1
	}
}
