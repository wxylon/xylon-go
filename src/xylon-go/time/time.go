package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 1)
	fmt.Println("start")
	go func() {
		//<-timer1.C
		fmt.Println(time.Now())
		fmt.Println(timer1.C)
	}()

	time.Sleep(10000)
}
