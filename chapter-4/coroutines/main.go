package main

import (
	"fmt"
	"tcc-coroutines/coroutines"

	"golang.org/x/tour/tree"
)

func main() {
	var find func(*coroutines.Coroutine, ...interface{}) []interface{}

	A := tree.New(2)
	B := tree.New(3)

	find = func(c *coroutines.Coroutine, args ...interface{}) []interface{} {
		t := (args[0]).(*tree.Tree)
		if t != nil {
			find(c, t.Left)
			c.Yield(t.Value)
			find(c, t.Right)
		}
		return nil
	}
	stepTree1 := coroutines.Create(find)
	stepTree2 := coroutines.Create(find)
	v1 := stepTree1.Resume(A)
	v2 := stepTree2.Resume(B)
	for len(v1) > 0 || len(v2) > 0 {
		if len(v1) > 0 && (len(v2) == 0 || v1[0].(int) < v2[0].(int)) {
			fmt.Printf("%v, ", v1[0])
			v1 = stepTree1.Resume()
		} else {
			fmt.Printf("%v, ", v2[0])
			v2 = stepTree2.Resume()
		}
	}
	fmt.Printf("\n")
}
