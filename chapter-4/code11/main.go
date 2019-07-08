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
	ID    int
	dirty bool
	Send  chan *fork
}

var forks []*fork
var gPhilosophers []*philosopher
var numPhilosophers int

func (p *philosopher) start() {
	fmt.Println("Starting philosopher", p.ID)
	for {
		// time thinking
		fmt.Println("Philosopher", p.ID, "is thinking")
		time.Sleep(time.Millisecond * time.Duration(rand.Int()%4000))
		fmt.Println("Philosopher", p.ID, "finished thinking")

		// must eat!
		fmt.Println("Philosopher", p.ID, "is trying to eat")
		select {
		case p.LeftFork.Send <- p.LeftFork:
			p.LeftFork = &fork{ID: 0}
		case p.RightFork.Send <- p.RightFork:
			p.RightFork = &fork{ID: 0}
		default:
		}
		for p.LeftFork.ID == 0 || p.RightFork.ID == 0 {
			select {
			case p.LeftFork.Send <- p.LeftFork:
				p.LeftFork = &fork{ID: 0}
			case p.RightFork.Send <- p.RightFork:
				p.RightFork = &fork{ID: 0}
			case fork := <-forks[p.ID-1].Send:
				p.LeftFork = fork
				p.LeftFork.Send = nil
			case fork := <-forks[p.ID%numPhilosophers].Send:
				p.RightFork = fork
				p.RightFork.Send = nil
			}
		}

		// fixed time eating
		fmt.Println("Philosopher", p.ID, "started eating")
		time.Sleep(time.Second * 1)
		fmt.Println("Philosopher", p.ID, "finished eating")
		p.LeftFork.Send = make(chan *fork)
		p.RightFork.Send = make(chan *fork)
		fmt.Printf("%+v, %+v, %+v, %+v, %+v\n", gPhilosophers[0], gPhilosophers[1], gPhilosophers[2], gPhilosophers[3], gPhilosophers[4])
		fmt.Printf("%+v, %+v, %+v, %+v, %+v\n", forks[0], forks[1], forks[2], forks[3], forks[4])
	}
}

func initForksAndPhilosophers(philosophers []*philosopher) {
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
		forks = append(forks, &fork{ID: i + 1, dirty: true, Send: make(chan *fork)})
		philosopher := &philosopher{
			ID:        i + 1,
			RightFork: &fork{},
			LeftFork:  &fork{},
		}
		philosophers = append(philosophers, philosopher)
	}
	gPhilosophers = philosophers
	initForksAndPhilosophers(philosophers)
	for _, p := range philosophers {
		go p.start()
	}
	time.Sleep(time.Minute * 5)
}
