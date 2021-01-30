package main

import "fmt"

func incrementor(start int) func() int {
	x := start

	return func() int {
		x++
		return x
	}
}

func main() {
	var incrementMe func() int = incrementor(5)

	fmt.Println(incrementMe())
	fmt.Println(incrementMe())
	fmt.Println(incrementMe())
	fmt.Println(incrementMe())
	fmt.Println(incrementMe())
	fmt.Println(incrementMe())
	fmt.Println(incrementMe())
	fmt.Println(incrementMe())
	/*
		6
		7
		8
		9
		10
		11
		12
		13
	*/
}
