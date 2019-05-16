package main

import (
	"fmt"
	"sync"
	"time"
)

var firstBarrier, secondBarrier sync.WaitGroup

func A(waitGroup *sync.WaitGroup) {
	fmt.Println("Starting A")
	time.Sleep(time.Second * 2)
	fmt.Println("A done. Waiting for the others!")
	waitGroup.Done()
	waitGroup.Wait()
	fmt.Println("Finished A")
}

func B(waitGroup *sync.WaitGroup) {
	fmt.Println("Starting B")
	time.Sleep(time.Millisecond * 300)
	fmt.Println("B done. Waiting for the others!")
	waitGroup.Done()
	waitGroup.Wait()
	fmt.Println("Finished B")
}

func C(waitGroup *sync.WaitGroup) {
	fmt.Println("Starting C")
	time.Sleep(time.Second * 5)
	fmt.Println("C done. Waiting for the others!")
	waitGroup.Done()
	waitGroup.Wait()
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
