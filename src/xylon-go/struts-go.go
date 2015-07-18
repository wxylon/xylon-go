package main

import "math"

type Point struct {
	x, y float64
}

func (self Point) length() float64 {
	return math.Sqrt(self.x*self.x + self.y*self.y)
}

func (self *Point) Scale(factor float64) {
	self.setX(self.x * factor)
	self.setY(self.y * factor)
}

func (self *Point) setX(x float64) {
	self.x = x
}

func (self *Point) setY(y float64) {
	self.y = y
}
