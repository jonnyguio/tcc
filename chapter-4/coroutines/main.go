package main

import (
	"fmt"
	"tcc-coroutines/coroutines"

	"golang.org/x/tour/tree"
)

func main() {
	var a, b int
	a = 1
	b = 1
	coro := coroutines.Create(func(c *coroutines.Coroutine, args ...interface{}) []interface{} {
		values := c.Yield(a, b)
		finalSum := a + b
		for _, value := range values {
			finalSum += value.(int)
		}
		return []interface{}{finalSum, 0}
	})
	values := coro.Resume()
	fmt.Println(values)
	values = coro.Resume(100)
	fmt.Println(values)

	A := tree.New(2)
	B := tree.New(3)
	fmt.Println(A)

	var find func(*coroutines.Coroutine, ...interface{}) []interface{}

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
			fmt.Println(v1)
			v1 = stepTree1.Resume()
		} else {
			fmt.Println(v2)
			v2 = stepTree2.Resume()
		}
	}
}
