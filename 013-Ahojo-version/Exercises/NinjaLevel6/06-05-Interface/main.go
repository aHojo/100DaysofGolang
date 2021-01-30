package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	sideLength int
}

func (s square) area() float64 {

	return float64(s.sideLength) * float64(s.sideLength)
}

type circle struct {
	radius int
}

func (c circle) area() float64 {

	area := math.Pi * math.Pow(float64(c.radius), float64(2))
	return area
}

func main() {
	circ := circle{
		radius: 5,
	}
	squar := square{
		sideLength: 5,
	}

	fmt.Println("Printing circle area")
	info(circ)
	fmt.Println("Printing square area")
	info(squar)

}

func info(s shape) {
	fmt.Println(s.area())
}
