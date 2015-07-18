package main

import (
	"flag"
	"fmt"
)

var (
	showVersion = flag.Bool("version", false, "print version string")
	httpAddress = flag.String("http-address", "0.0.0.0:4151", "<addr>:<port> to listen on for HTTP clients")
)

func main() {
	fmt.Printf("%s\n", "hello, world")
	printNumber()

	// flag.Parse()
	// fmt.Println(*showVersion)
	// fmt.Println(*httpAddress)

}

func printNumber() {
	var count int
	count = 10

	sum := 0

	for i := 0; i < count; i++ {
		sum = sum + i
		fmt.Printf("i = %d, sum = %d\n", i, sum)
	}
}
