package main

import "fmt"

func main() {

	sumVariadic := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	sumSlice := sum1([]int{10, 11, 12, 13, 14, 15})

	fmt.Println(sumVariadic)
	fmt.Println(sumSlice)
}

func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}

func sum1(nums []int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}
