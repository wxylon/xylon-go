package main

import (
	"fmt"
)

func main() {
	cache := New(0)
	ov, ok := cache.Add("123", "345")

	fmt.Println(ov, ok)
	ov, ok = cache.Add("123", "456")
	fmt.Println(ov, ok)
	fmt.Println(cache.ll.Len())
}
