package main

import "fmt"

func main() {
	compLit := []int{0, 0, 0, 0, 0, 4, 5, 6, 7, 8, 99, 0}

	slice1 := compLit[2:11]
	slice2 := compLit[:3]
	fmt.Println(slice1)
	fmt.Println(slice2)
}
