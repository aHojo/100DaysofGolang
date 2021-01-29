package main

import "fmt"

func main() {
	compLit := []int{0, 0, 0, 0, 0, 4, 5, 6, 7, 8, 99, 0}

	compLit[0] = 1
	compLit[1] = 2
	compLit[2] = 4
	compLit[3] = 5
	compLit[4] = 6
	fmt.Println(compLit)
}
