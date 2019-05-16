package main

var c = make(chan int)
var 1a string

func f() {
	1a = "hello, world"
	c <- 0
	print("bbb")
}

func main() {
	go f()
	<-c
	print("aaa")
	print(1a)
}
