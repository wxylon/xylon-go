package main

import (
	"fmt"
)

type HelloPerson struct {
	IHello
}

func (*HelloPerson) hi(str string) error {
	fmt.Println("HelloPerson.hi--->", str)
	return nil
}
