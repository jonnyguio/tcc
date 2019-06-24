package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type philosopher struct {
	ID                  int
	LeftFork, RightFork *fork
}

type fork struct {
	ID      int
	dirty   bool
	Request chan *philosopher
}

var forks []*fork
var gPhilosophers []*philosopher
var numPhilosophers int

func (p *philosopher) start() {
	fmt.Println("Starting philosopher", p.ID)
	useCh := make(chan bool)
	go func() {
		canUse := false
		for {
			select {
			case canUse = <-useCh:
			default:
				if !canUse {
					if p.LeftFork != nil && p.LeftFork.dirty {
						select {
						case requester := <-p.LeftFork.Request:
							if p != requester {
								fmt.Println("Philosopher", p.ID, "gave left fork to philosopher", requester.ID)
								p.LeftFork = nil
							}
						default:
						}
					}
					if p.RightFork != nil && p.RightFork.dirty {
						select {
						case requester := <-p.RightFork.Request:
							if p != requester {
								fmt.Println("Philosopher", p.ID, "gave right fork to philosopher", requester.ID)
								p.RightFork = nil
							}
						default:
						}
					}
				}
			}
		}
	}()
	for {
		// time thinking
		fmt.Println("Philosopher", p.ID, "is thinking")
		time.Sleep(time.Millisecond * time.Duration(rand.Int()%4000))
		fmt.Println("Philosopher", p.ID, "finished thinking")

		// must eat!
		for {
			fmt.Println("Philosopher", p.ID, "is trying to eat")
			useCh <- true
			if p.LeftFork != nil && p.RightFork != nil {
				break
			}
			useCh <- false
			select {
			case forks[p.ID-1].Request <- p:
				p.LeftFork = forks[p.ID-1]
				p.LeftFork.dirty = false
			case forks[(p.ID)%numPhilosophers].Request <- p:
				p.RightFork = forks[(p.ID)%numPhilosophers]
				p.RightFork.dirty = false
			}
		}

		// fixed time eating
		fmt.Println("Philosopher", p.ID, "started eating")
		time.Sleep(time.Second * 1)
		fmt.Println("Philosopher", p.ID, "finished eating")
		p.LeftFork.dirty, p.RightFork.dirty = true, true
		useCh <- false
	}
}

func initForks(philosophers []*philosopher) {
	for k, v := range forks {
		if k%2 == 1 {
			philosophers[k-1].RightFork = v
		} else {
			philosophers[k].LeftFork = v
		}
	}
}

func main() {
	var err error
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <number of philosophers>\n", os.Args[0])
		os.Exit(1)
	}
	numPhilosophers, err = strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	philosophers := []*philosopher{}
	for i := 0; i < numPhilosophers; i++ {
		forks = append(forks, &fork{ID: i + 1, dirty: true, Request: make(chan *philosopher)})
		philosopher := &philosopher{
			ID: i + 1,
		}
		philosophers = append(philosophers, philosopher)
	}
	initForks(philosophers)
	gPhilosophers = philosophers
	for _, p := range philosophers {
		go p.start()
	}
	time.Sleep(time.Minute * 5)
}
