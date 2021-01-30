package main

import "fmt"

func main() {

	sumVariadic := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)

	fmt.Println(sumVariadic)

}

func sum(nums ...int) int {
	defer sum1(nums)
	fmt.Println("I'm running in sum")
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}

func sum1(nums []int) int {
	fmt.Println("I'm running in sum1")
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}
