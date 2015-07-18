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

type B struct {
	A
}

func NewB(param int) *B {
	return &B{A{param}}
}

func NewB0(param int) B {
	return B{A{param}}
}

func (this *B) Read(buf *byte, cb int) int {
	return cb * cb
}

func (this *B) Write(bug *byte, cb int) int {
	return cb * cb
}

func (this *B) Foo() {
	fmt.Println("B.Foo:", this.a)
}

func main() {

	/**接口指向父实现类 A,调用父类的方法*/
	var pA IReadWirter = NewA(8)
	fmt.Println(pA.Read(nil, 10))
	fmt.Println(pA.Write(nil, 11))
	fmt.Println(&pA)
	fmt.Println("===========================")

	/**接口指向子实现类 B,调用子类的方法*/
	var pB IReadWirter = NewB(9)
	fmt.Println(pB.Read(nil, 10))
	fmt.Println(pB.Write(nil, 11))
	fmt.Println(&pB)
	fmt.Println("===========================")

	/**实例化A*/
	oA := new(A)
	oA.a = 100
	fmt.Println(oA.Read(nil, 10))
	fmt.Println(oA.Write(nil, 11))
	fmt.Println(&oA)

	/**实例化B*/
	fmt.Println("===========================")
	oB := new(B)
	oB.a = 100
	fmt.Println(oB.Read(nil, 10))
	oB.Foo()
	fmt.Println(oB.Write(nil, 11))
	oB.Foo()
	fmt.Println(&oB)

	fmt.Println("===========================")
	/**接口指向子实现类 B,调用子类的方法*/
	mB := NewB0(1000)
	mB.a = 100
	fmt.Println(mB.Read(nil, 10))
	fmt.Println(mB.Write(nil, 11))
	fmt.Println(mB)
	mB.Foo()
}
