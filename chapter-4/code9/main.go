package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func Step(a, b, c [][]int, row, col, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Println(row, col)
	// fmt.Println(len(a), len(b), len(c))
	for i, j := 0, 0; i < n; i, j = i+1, j+1 {
		c[row][col] += a[row][i] * b[j][col]
	}
}

func MultiplyMatrix(a, b, c [][]int, n int) {
	wg := &sync.WaitGroup{}
	// fmt.Println(len(a), len(b), len(c))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			wg.Add(1)
			go Step(a, b, c, i, j, n, wg)
		}
	}
	wg.Wait()
}

func initMatrix(a *[][]int, n, init int) {
	*a = make([][]int, n)
	for i := 0; i < n; i++ {
		(*a)[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if init != 0 {
				(*a)[i][j] = rand.Int() % 100
			}
		}
	}
}

func printMatrix(a [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	var A, C [][]int
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <size of squared matrix>\n", os.Args[0])
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Failed to convert matrix size as number")
		panic(err)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	initMatrix(&A, n, 1)
	initMatrix(&C, n, 0)
	// fmt.Println(len(A), len(A), len(C))

	//fmt.Println("Matrix A:")
	//printMatrix(A, n)

	//fmt.Println("Matrix C:")
	//printMatrix(C, n)
	start := time.Now()
	MultiplyMatrix(A, A, C, n)
	end := time.Now().Sub(start)
	//fmt.Println("Result Matric C:")
	//printMatrix(C, n)
	fmt.Println(end.String())
}
