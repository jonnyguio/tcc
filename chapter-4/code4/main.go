package main

import (
	"fmt"
	"sync"
	"time"
)

var firstBarrier, secondBarrier sync.WaitGroup

func A(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println("Starting A")
	time.Sleep(time.Second * 2)
	fmt.Println("Finished A")
}

func B(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println("Starting B")
	time.Sleep(time.Millisecond * 300)
	fmt.Println("Finished B")
}

func C(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println("Starting C")
	time.Sleep(time.Second * 5)
	fmt.Println("Finished C")
}

func main() {
	fmt.Println("Start")
	firstBarrier.Add(2)
	secondBarrier.Add(4)
	go A(&firstBarrier)
	go B(&firstBarrier)
	go C(&secondBarrier)

	fmt.Println("Waiting A and B")
	firstBarrier.Wait()
	fmt.Println("Finished waiting A and B")

	go A(&secondBarrier)
	go B(&secondBarrier)
	go C(&secondBarrier)

	fmt.Println("Waiting 2C, A and B")
	secondBarrier.Wait()
	fmt.Println("Finished waiting 2C, A and B")

	fmt.Println("Finish")
}
