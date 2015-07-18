//http://outofmemory.cn/code-snippet/1172/Go-language-li-For-do-statement
package main

import (
	"fmt"
)

const (
	count int = 10
)

func main() {
	for0()
	for1()
	for2()
	for3()
}

func for0() {
	sum := 0
	for i := 0; i < count; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func for1() {
	i, sum := 0, 0
	for i < count {
		sum += i
		i++
	}
	fmt.Println(sum)
}

func for2() {
	i, sum := 0, 0
	for {
		if i >= count {
			break
		}
		sum += i
		i++
	}
	fmt.Println(sum)
}

func for3() {
	i, sum := 0, 0
	for {
		if i >= count {
			break
		}
		sum += i
		i++
	}
	fmt.Println(sum)
}
