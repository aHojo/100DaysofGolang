package main

import "fmt"

func main() {

	const (
		c1 = iota
		c2
		c3
	)

	fmt.Println(c1, c2, c3)

	const (
		NORTH = iota
		EAST
		SOUTH
		WEST
	)
	fmt.Println(NORTH, EAST, WEST, SOUTH)

	const (
		a = iota * 2 //0
		b            // 2
		c            // 4
	)
	fmt.Println(a, b, c)

	const (
		d = (iota * 2) + 1 //0
		_                  // skipped
		f                  // 5
	)

	fmt.Println(d, f)
}
