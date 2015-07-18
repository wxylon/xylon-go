//http://blog.csdn.net/xushiweizh/article/details/7034346
package main

import "fmt"

type IReadWirter interface {
	Read(buf *byte, cb int) int
	Write(buf *byte, cb int) int
}

type A struct {
	a int
}

func NewA(param int) *A {
	return &A{a: param}
}

func (this *A) Read(buf *byte, cb int) int {
	return cb
}

func (this *A) Write(bug *byte, cb int) int {
	return cb
}

func main() {
	var pA *A = NewA(8)
	// pA.Write(nil, 11)
	fmt.Println(pA)
	fmt.Println(pA.Read(nil, 11))
}
