package main

import "fmt"

func main() {
	var p *[10]int = new([10]int)
	fmt.Println(&p)
	testArray(p)

	var p1 []int = make([]int, 10)
	fmt.Println(&p1)
	testArray1(p1)
}

func testArray(p *[10]int){
	fmt.Println(&p)
}

func testArray1(p1 []int){
	fmt.Println(p1)
}
