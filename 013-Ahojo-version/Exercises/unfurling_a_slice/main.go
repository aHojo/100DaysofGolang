package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}

	x := sum(xi...) // This is where we expand the slice, or aka unfurl

	fmt.Println("The total is", x)

}

func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is", sum)
	}

	return sum
}
