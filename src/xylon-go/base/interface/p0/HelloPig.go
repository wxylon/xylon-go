package main

import (
	"fmt"
)

type HelloPig struct {
	IHello
}

func (*HelloPig) hi(str string) error {
	fmt.Println("HelloPig.hi--->", str)
	return nil
}
