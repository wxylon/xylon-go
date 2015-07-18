package main

import (
	"flag"
	"fmt"
	"os"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")
var c byte = 1

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.Parse()
	var s string = ""
	fmt.Print("参数个数：", flag.NArg(), "; values:")
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if !*omitNewline {
		s += Newline
	}
	os.Stdout.WriteString(s)
}
