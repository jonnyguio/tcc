package main

import (
	"fmt"
	"time"
)

func processData() {
	time.Sleep(10 * time.Second)
}

func main() {
	go func() {
		for {
			fmt.Println("working..")
			time.Sleep(2 * time.Second)
		}
	}()
	processData()
}
