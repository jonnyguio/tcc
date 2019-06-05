package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	A, C := MatrixInt{}, MatrixInt{}
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
	// fmt.Println(len(A), len(A), len(C))

	A.InitSquared(n, 1000, true)
	C.InitSquared(n, 0, false)

	/*fmt.Println("Matrix A:")
	A.Print()
	fmt.Println("Matrix C:")
	C.Print()*/

	start := time.Now()
	A.Multiply(A, C, parallel)
	end := time.Now().Sub(start)

	/*fmt.Println("Result Matric C:")
	C.Print()*/
	fmt.Println(end.String())
}
