package main

import (
	"fmt"
	"math"
)

type Figure struct {
	length int
	width  int
}

func (f Figure) Area() float64 {
	var area float64

	if f.width == 0 {
		area = math.Pow(float64(f.length), 2) / (4 * math.Pi)
	} else {
		area = float64(f.length * f.width)
	}
	return area
}

func (f Figure) Type() {
	if f.width == 0 {
		fmt.Println("It's Rectangle")
	} else {
		fmt.Println("It's Circle")
	}
}

func main() {
	Rectangle := Figure{3, 5}
	Circle := Figure{9, 0}

	Rectangle.Type()
	fmt.Println(Rectangle.Area())

	Circle.Type()
	fmt.Println(Circle.Area())
}
