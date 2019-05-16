package main

import (
	"fmt"
	"time"
)

func doA(ch chan int) {
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("doA")
	ch <- 'A'
	fmt.Println("finishes doA")
}

func doB(ch chan int) {
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("doB")
	ch <- 'B'
	fmt.Println("finishes doB")
}

func main() {
	chA := make(chan int, 1)
	chB := make(chan int, 1)
	lastA := time.Now()
	lastB := time.Now()
	for {
		if time.Since(lastA) > time.Second * 5 {
			go doA(chA)
		}
		if time.Since(lastB) > time.Second * 5 {
			go doB(chB)
		} 
		time.Sleep(time.Second * 4)
		select {
			case v1 := <-chA:
				fmt.Println(v1)
				lastA = time.Now()
			case v2 := <-chB:
				fmt.Println(v2)
				lastB = time.Now()
			default:
				fmt.Println("none to send")
				time.Sleep(time.Second * 1)
		}
	}
	fmt.Println("finishes")
}

