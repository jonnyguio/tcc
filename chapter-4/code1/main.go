package main

import (
	"math/rand"
	"time"
	"sync"
)


func main() {
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UTC().UnixNano())
	streamSize := 32
	stream := make(chan []byte)

	wg.Add(2)
	go consumer(stream, streamSize, wg)
	go producer(stream, streamSize, wg)
	wg.Wait()
}

