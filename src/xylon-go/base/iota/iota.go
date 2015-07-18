/**
sources:http://blog.163.com/huv520@126/blog/static/2776523920100313211221/
notes:
	1.iota关键字用在const常量中;
	2.未指定常量值情况，选取最近的表达式
	3.iota作用域未括号或者const关键字范围
*/
package main

import "fmt"

const (
	a1 = 1 << iota // 1 << 0
	a2             // 1 << 1
	a3 = iota      // 2
)

const (
	b1 = 1 << iota // a == 1 (iota has been reset)    1*2^0
	b2 = 1 << iota // b == 2                          1*2^1
	b3 = 1 << iota // c == 4                          1*2^2
)

const ( // iota is reset to 0
	c1 = iota // c0 == 0
	c2 = iota // c1 == 1
	c3 = iota // c2 == 2
)

const d1 = iota // d1 == 0 (iota has been reset)
const d2 = iota // d2 == 0 (iota has been reset)

const (
	e1, f1 = 1 << iota, 1<<iota - 1 // e1 == 1, f1 == 0
	e2, f2                          // e2 == 2, f2 == 1
	_, _                            // skips iota == 2
	e4, f4                          // e4 == 8, f4 == 7
)

func main() {
	fmt.Println("a1=", a1, "; a2=", a2, "; a3=", a3)
	fmt.Println("b1=", b1, "; b2=", b2, "; b3=", b3)
	fmt.Println("c1=", c1, "; c2=", c2, "; c3=", c3)
	fmt.Println("d1=", d1, "; d2=", d2)
	fmt.Println("e1=", e1, "; e2=", e2, "; e4=", e4)
	fmt.Println("f1=", f1, "; f2=", f2, "; f4=", f4)

	// a1= 1 ; a2= 2 ; a3= 2
	// b1= 1 ; b2= 2 ; b3= 4
	// c1= 0 ; c2= 1 ; c3= 2
	// d1= 0 ; d2= 0
	// e1= 1 ; e2= 2 ; e4= 8
	// f1= 0 ; f2= 1 ; f4= 7
}
