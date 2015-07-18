package main

import (
	"fmt"
)

var a string
var sid = make(chan int, 1)

func f() {
	fmt.Println(a)
	sid <- 1
	//print(a)
}

func hello() {
	a = "hello, world"
	go f()
}

func main() {
	hello()
	<-sid
}
