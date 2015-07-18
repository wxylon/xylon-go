package main

import (
	"fmt"
	"time"
)

type Test struct {
	waitGroup *WaitGroupWrapper
	Name      string
}

func NewTest(name string, w *WaitGroupWrapper) *Test {
	test := &Test{Name: name, waitGroup: w}
	test.waitGroup.Wrap(func() { test.router() })
	return test
}

func (t *Test) router() {
	for {
		fmt.Println("Test: doing ...", t.Name)
	}
}

func (t *Test) string1() string {
	return string(t.Name)
}

func main() {
	w := &WaitGroupWrapper{}
	t := NewTest("123", w)
	t.string1()
	time.Sleep(10 * 1000)
}
