package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)

func producer(stream chan []byte, streamSize int, wg sync.WaitGroup) {
	randomInput := make([]byte, streamSize)
	for {
		n, err := rand.Read(randomInput)
		if err != nil || n != streamSize {
			fmt.Println("Error creating randomInput, exiting producer.")
			break
		}
		fmt.Println("Producer writing randomInput")
		stream <- randomInput
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}
