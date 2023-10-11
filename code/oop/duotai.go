package main

import (
	"fmt"
	"math"
)

// 通过接口实现多态

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}

	fmt.Println(GetArea(r)) // 输出：12
	fmt.Println(GetArea(c)) // 输出：78.53981633974483
}
