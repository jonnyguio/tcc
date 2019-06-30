package main

import (
	"crypto/rand"
	"fmt"
	"sync"
)

type Writer struct {
	ID      int
	Writers *sync.Mutex
}

func (w *Writer) start(stream []byte, size int) {
	for {
		w.Writers.Lock()

		fmt.Printf("Writer %d wrinting byte string %v\n", w.ID, stream)
		rand.Read(stream)

		w.Writers.Unlock()
	}
}
