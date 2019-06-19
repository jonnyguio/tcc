package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	A, B, C := MatrixInt{}, MatrixInt{}, MatrixInt{}
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <size of squared matrix> <parallel: true or false>\n", os.Args[0])
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Failed to convert matrix size as number")
		panic(err)
	}
	parallel, err := strconv.ParseBool(os.Args[2])
	if err != nil {
		fmt.Println("Failed to convert option parallel")
		panic(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	A.InitSquared(n, 1000, true)
	B.InitSquared(n, 1000, true)
	C.InitSquared(n, 0, false)

	start := time.Now()
	A.Multiply(B, C, parallel)
	end := time.Now().Sub(start)

	fmt.Println(end.String())
}
