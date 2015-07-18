package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func Sum(values []int, results chan chan int) {
	result := make(chan int, 1)
	results <- result
	fmt.Println(values)
	time.Sleep(time.Second * 5)
	sum := 0
	for _, value := range values {
		sum += value
	}
	result <- sum
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	values1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	results := make(chan chan int, 1)
	go Sum(values, results)
	//fmt.Println(<-result)
	//sum1 := <-result,
	go Sum(values1, results)

	chan_slice := make([]chan int, 0)
	var count int
	for {
		select {

		case single_chan := <-results:
			count++
			log.Println("the single_chan is ", single_chan)
			chan_slice = append(chan_slice, single_chan)
			if count == 2 {
				runtime.Goexit()
			}
		}
	}

	// sum1, sum2 := <-result, <-result
	// fmt.Println("Result:", sum1, sum2, sum1+sum2)
	//time.Sleep(time.Hour * 1)
}
