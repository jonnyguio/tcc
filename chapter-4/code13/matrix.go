package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type MatrixInt [][]int

func (a MatrixInt) Step(b, c MatrixInt, i, k int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < len(a); j++ {
		c[i][j] += a[i][k] * b[k][j]
	}
}

func (a MatrixInt) Multiply(b, c MatrixInt, parallel bool) {
	wg := &sync.WaitGroup{}
	for i := 0; i < len(a); i++ {
		for k := 0; k < len(a); k++ {
			wg.Add(1)
			if parallel {
				go a.Step(b, c, i, k, wg)
			} else {
				a.Step(b, c, i, k, wg)
			}
		}
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
