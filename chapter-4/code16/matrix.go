package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type MatrixInt [][]int

func (c MatrixInt) Step(a, b MatrixInt, id, nthreads int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := id; i < len(a); i += nthreads {
		for k := 0; k < len(a[0]); k++ {
			for j := 0; j < len(a[0]); j++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
}

func (c MatrixInt) Multiply(a, b MatrixInt, threads int) {
	wg := &sync.WaitGroup{}
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go c.Step(a, b, i, threads, wg)
	}
	wg.Wait()
}

func (a *MatrixInt) InitSquared(newSize, maxCell int, fillRandomly bool) {
	*a = make([][]int, newSize)
	for i := 0; i < newSize; i++ {
		(*a)[i] = make([]int, newSize)
		if fillRandomly {
			for j := 0; j < newSize; j++ {
				(*a)[i][j] = rand.Int() % maxCell
			}
		}
	}

}

func (a MatrixInt) Print() {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			fmt.Printf("%d ", (a)[i][j])
		}
		fmt.Println("")
	}
}
