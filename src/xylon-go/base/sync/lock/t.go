//http://www.dotcoo.com/golang-sync-mutex
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
	t1(3)
	fmt.Printf("%s\n", "exit!")
}

func t1(i int) {
	fmt.Println(i, "--->lock start")
	lock.Lock()
	fmt.Println(i, "--->lock")
	time.Sleep(time.Second * 10)
	lock.Unlock()
	fmt.Println(i, "--->unlock")
}
