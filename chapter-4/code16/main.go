package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	A, B, C := MatrixInt{}, MatrixInt{}, MatrixInt{}
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

	A.InitSquared(n, 1000, true)
	B.InitSquared(n, 1000, true)
	C.InitSquared(n, 0, false)

	start := time.Now()
	A.Multiply(B, C, runtime.GOMAXPROCS(0))
	end := time.Now().Sub(start)

	fmt.Println(end.String())
}
