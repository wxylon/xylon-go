package main

import "fmt"

func main() {
	var p *Point = new(Point)
	p.setX(200)
	p.setY(300)

	fmt.Println(p.length())
	fmt.Println(p.x, p.y)
	p.length()
	fmt.Printf("Hello, world; or Καλημέρα κόσμε; or こんにちは 世界\n")
}
