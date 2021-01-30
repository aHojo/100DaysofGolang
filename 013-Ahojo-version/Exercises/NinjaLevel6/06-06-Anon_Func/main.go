package main

import "fmt"

type Num struct {
	toGetTripled int
}

func main() {

	num := Num{}
	func(x int, num *Num) {

		num.toGetTripled = x * 3
	}(5, &num)

	fmt.Println(num.toGetTripled)
}
