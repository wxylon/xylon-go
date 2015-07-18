package main

import (
	"fmt"
)

type BigStruct struct {
	C01 uint64
}

func Invoke1(a *BigStruct) uint64 {
	a.C01++
	return a.C01
}

func Invoke2(a BigStruct) uint64 {
	a.C01++
	return a.C01
}

func (a *BigStruct) Invoke3() uint64 {
	a.C01++
	return a.C01
}

func (a BigStruct) Invoke4() uint64 {
	a.C01++
	return a.C01
}

func main() {
	var a = new(BigStruct)
	for i := 0; i < 3; i++ {
		fmt.Println("指针传递：", Invoke1(a))
	}

	var b = BigStruct{}
	for i := 0; i < 3; i++ {
		fmt.Println("值传递：", Invoke2(b))
	}

	var c = BigStruct{}
	for i := 0; i < 3; i++ {
		fmt.Println("指针传递：", c.Invoke3())
	}

	var d = BigStruct{}
	for i := 0; i < 3; i++ {
		fmt.Println("值传递：", d.Invoke4())
	}
}
