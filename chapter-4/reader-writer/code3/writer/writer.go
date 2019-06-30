package writer

import (
	"crypto/rand"
	"fmt"
	"sync"
)

type Writer struct {
	ID      int
	Writers *sync.Mutex
	Readers *sync.Mutex
}

var local, count = &sync.Mutex{}, 0

func (w *Writer) Start(stream []byte, size int) {
	for {
		local.Lock()
		count++
		if count == 1 {
			w.Readers.Lock()
		}
		local.Unlock()

		w.Writers.Lock()
		fmt.Printf("Writer %d wrinting byte string %v\n", w.ID, stream)
		rand.Read(stream)
		w.Writers.Unlock()

		local.Lock()
		count--
		if count == 1 {
			w.Readers.Unlock()
		}
		local.Unlock()
	}
}
