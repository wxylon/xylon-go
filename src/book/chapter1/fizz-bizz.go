package main

import "fmt"

const(
	fizz = 3
	bizz = 5
	count = 100
)
func main(){
	var p bool
	for i:=1; i<=count;i++{
		p = false
		if i% fizz == 0{
			fmt.Print("Fizz")
			p = true
		}
		
		if i % bizz == 0{
			fmt.Print("Buzz")
			p = true
		} 
		
		if !p {
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}