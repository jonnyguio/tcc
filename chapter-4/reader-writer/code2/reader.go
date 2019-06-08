package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/semaphore"
)

type Reader struct {
	ID      int
	Readers *semaphore.Weighted
	Writers *semaphore.Weighted
}

func (r *Reader) start(stream []byte, size int) {
	for {
		r.Writers.Acquire(context.Background(), 1)
		r.Readers.Acquire(context.Background(), 1)
		r.Writers.Release(1)
		fmt.Printf("Reader %d read byte string %v\n", r.ID, stream)
		r.Readers.Release(1)
	}
}
