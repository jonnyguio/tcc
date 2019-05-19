package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type Philosopher struct {
	ID                  int
	Waiter              *sync.Mutex
	LeftFork, RightFork *Fork
}

type Fork struct {
	available bool
}

func (p *Philosopher) start() {
	fmt.Println("Starting philosopher ", p.ID)
	for {
		// time thinking
		fmt.Println("Philosopher", p.ID, "is thinking...")
		time.Sleep(time.Millisecond * time.Duration(rand.Int()%4000))
		fmt.Println("Philosopher", p.ID, "finished thinking!")

		// must eat!
		fmt.Println("Philosopher", p.ID, "checking if he can get his forks!")
		for {
			p.Waiter.Lock()
			if p.LeftFork.available && p.RightFork.available {
				p.LeftFork.available = false
				p.RightFork.available = false
				p.Waiter.Unlock()
				fmt.Println("Philosopher", p.ID, "got his forks")
				break
			}
			p.Waiter.Unlock()
		}

		// fixed time eating
		fmt.Println("Philosopher", p.ID, "is eating...")
		time.Sleep(time.Second * 1)
		p.LeftFork.available = true
		p.RightFork.available = true
		fmt.Println("Philosopher", p.ID, "dropped his forks...")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <number of philosophers>\n", os.Args[0])
		os.Exit(1)
	}
	numPhilosophers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Failed to convert matrix size as number")
		panic(err)
	}

	var nextFork, firstFork, lastFork *Fork
	waiterMutex := &sync.Mutex{}

	firstFork = &Fork{available: true}
	lastFork = firstFork

	for i := 0; i < numPhilosophers; i++ {
		if i == numPhilosophers-1 {
			nextFork = firstFork
		} else {
			nextFork = &Fork{available: true}
		}
		philosopher := &Philosopher{
			ID:        i + 1,
			Waiter:    waiterMutex,
			LeftFork:  lastFork,
			RightFork: nextFork,
		}
		lastFork = nextFork
		fmt.Printf("%d %p %p\n", philosopher.ID, philosopher.LeftFork, philosopher.RightFork)
		go philosopher.start()
	}
	time.Sleep(time.Minute * 5)
}
