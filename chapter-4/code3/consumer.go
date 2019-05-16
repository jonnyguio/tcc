package main

import (
	"fmt"
	"sync"
	"time"
)

type Consumer struct {
	ID         int
	stream     chan []byte
	streamSize int
}

func (c *Consumer) start(stream chan []byte, streamSize int, wg *sync.WaitGroup) {
	c.stream = stream
	c.streamSize = streamSize
	streamBytes := make([]byte, c.streamSize)
	ok := true
	for {
		if streamBytes, ok = <-c.stream; ok {
			fmt.Printf("Consumer %d read from stream %d bytes!\n%v\n", c.ID, c.streamSize, streamBytes)
		} else {
			break
		}
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
