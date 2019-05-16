package main

import (
	"fmt"
	"sync"
	"time"
)

func consumer(stream chan []byte, streamSize int, wg sync.WaitGroup) {
	streamBytes := make([]byte, streamSize)
	ok := true
	for {
		if streamBytes, ok = <-stream; ok {
			fmt.Printf("Read from stream %d bytes!\n%v\n", streamSize, streamBytes)
		} else {
			break
		}
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
