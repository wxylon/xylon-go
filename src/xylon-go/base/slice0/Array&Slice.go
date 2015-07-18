//http://kejibo.com/golang-slice-array-reflect/

package main

import (
	"fmt"
	"reflect"
)

func sliceType() {
	vs := []interface{}{
		[]int{1, 2, 3}, //slice 切片
		[]int{1, 2, 3}, //切片再切还是切片
		//[]int{1, 2, 3}[:],  
		[3]int{1, 2},    //array 数组，确定数组长度
		[3]int{1, 2, 3}, //数组切一下，切出个 slice切片
		//[3]int{1, 2, 3}[:],
		[...]int{1, 2, 3}, //array 数组，由编译器自动计算数组长度。
		make([]int, 5, 5),
		new([]int),
	}
	fmt.Println(vs, reflect.ValueOf(vs).Kind())
	for i, v := range vs {
		rv := reflect.ValueOf(v) //进入疯狂的reflect世界
		fmt.Println(i, rv.Kind())
	}
}

//http://segmentfault.com/q/1010000000122913
func ArrayParams(a [3]int) [3]int {
	a[0] = 10
	return a
}

func SliceParams(a []int) []int {
	a[0] = 10
	return a
}

func run() {
	a := [...]int{1, 2, 3}
	b := ArrayParams(a)
	fmt.Println("a:", a, "   b:", b)
	fmt.Println("a=b:", (a == b))
	fmt.Println("========================================")
	c := []int{1, 2, 3}
	d := SliceParams(c)
	fmt.Println("c:", c, "   d:", d)
	//fmt.Println("c=d:", (c == d))
	fmt.Println("&c:", &c, "&d:", &d)
}

func main() {
	//sliceType()
	run()
}
