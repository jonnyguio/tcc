package main

import (
	"context"
	"crypto/rand"
	"fmt"

	"golang.org/x/sync/semaphore"
)

type Writer struct {
	ID      int
	Readers *semaphore.Weighted
	Writers *semaphore.Weighted
}

var internalWriter = semaphore.NewWeighted(1)

func (w *Writer) start(stream []byte, size int) {
	for {
		internalWriter.Acquire(context.Background(), 1)
		w.Readers.Wait()
		w.Writers.Acquire(context.Background(), 1)

		fmt.Printf("Writer %d wrinting byte string %v\n", w.ID, stream)
		rand.Read(stream)

		w.Writers.Release(1)
		internalWriter.Release(1)
	}
}
