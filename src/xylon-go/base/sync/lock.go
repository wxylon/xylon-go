package main

import (
	"fmt"
	"sync"
	"time"
)

var lock *sync.Mutex

func main() {
	lock = new(sync.Mutex)
	go t1(2)
	go t1(3)
	fmt.Printf("%s\n", "exit!")
}

func t1(i int) {
	println(i, "--->lock start")
	lock.Lock()
	println(i, "--->lock")
	time.Sleep(time.Second * 10)
	lock.Unlock()
	println(i, "--->unlock")
}
