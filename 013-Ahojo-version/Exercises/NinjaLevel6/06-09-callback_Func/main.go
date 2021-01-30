package main

import "fmt"

func multiplyByTwo(ints []int) []int {

	y := []int{}
	for _, v := range ints {
		y = append(y, v*2)
	}

	return y
}

func doWork(f func([]int) []int, x []int) {

	fmt.Println("Doing the work: ", f(x))
}

func main() {

	z := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Z before", z)
	doWork(multiplyByTwo, z)
	fmt.Println("Z After: ", z)

}
