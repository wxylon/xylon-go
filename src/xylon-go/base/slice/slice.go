package main

import (
	"fmt"
	"reflect"
)

//http://blog.golang.org/2011/01/go-slices-usage-and-internals.html

func init0() {
	/**数组不需要显式初始化;零值的数组是一个随时可用的数组,该数组的元素本身零*/
	var a [4]int
	fmt.Println(a)
	a[0] = 1
	i := a[0]
	fmt.Println(i == 1)
	fmt.Println(a[2] == 0)
}

func init1() {
	b := [3]string{"m", "f"}
	c := [...]string{"m", "f"}
	d := []string{"m", "f"}
	fmt.Println(b, reflect.ValueOf(b).Kind(), len(b), cap(b))
	fmt.Println(c, reflect.ValueOf(c).Kind(), len(c), cap(c))
	fmt.Println(d, reflect.ValueOf(d).Kind(), len(d), cap(d))
}

//make 动态创建指定的切片类型。
func init2() {
	var s0 [5]byte
	s1 := make([]byte, 5, 5)
	fmt.Println(s1, len(s1), cap(s1), reflect.ValueOf(s1).Kind())
	//[0 0 0 0 0] 5 5
	s1 = append(s1, '9')
	fmt.Println(s1, len(s1), cap(s1), reflect.ValueOf(s1).Kind())
	//[0 0 0 0 0 57] 6 10
	fmt.Println(s0, reflect.ValueOf(s0).Kind())

}

func main() {
	//init0()
	//init1()
	init2()
	//var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(a)
	//a = append(a, 0)
	//fmt.Println(a)
}
