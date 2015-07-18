package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(x, n int) {
			for y := 0; y < n; y++ {
				fmt.Printf("[%d]: %d\n", x, y)
			}
			wg.Done()
		}(i, rand.Intn(5))
	}
	wg.Wait()
}
