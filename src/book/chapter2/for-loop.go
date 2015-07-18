package main

import (
	"fmt"
	"time"
)

const (
	count int = 10
)

func main() {
	t0()
	t1()
}

func t0() {
	t := time.Now().Nanosecond()
	fmt.Println(time.Now().Nanosecond() - t)
	i := 0
I:
	fmt.Printf("%d,", i)
	i++
	if i < count {
		goto I
	}
	fmt.Println(time.Now().Nanosecond() - t)
}

func t1() {
	t := time.Now().Nanosecond()
	for i := 0; i < count; i++ {
		fmt.Printf("%d\n", i)
	}
	fmt.Println(time.Now().Nanosecond() - t)
}
