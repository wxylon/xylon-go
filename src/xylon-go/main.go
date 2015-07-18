package main

import "fmt"

func main() {
	// 简单声明
	var sum int
	sum = 100
	// 声明并初始化
	//var num int = 100

	var a, b *int
	a = &sum
	b = a

	//声明并初始化
	var label string = "ss100"
	//声明并初始化
	name := "sfd"

	//赋值操作
	label = "2323242"

	fmt.Println(a, b, label, name)

	fmt.Println("Hello World!", sum, sum+sum)

	if result := add(); result > 0 {
		fmt.Println("大于0")
	} else {
		fmt.Println("小于0")
	}

	switchtest := 'k'

	switch switchtest {
	case 'a', 'b':
		fmt.Println("a || b")
	case 'c':
		fmt.Println("c")
	default:
		fmt.Println("k")
	}

}

func add() int {
	return 0
}
