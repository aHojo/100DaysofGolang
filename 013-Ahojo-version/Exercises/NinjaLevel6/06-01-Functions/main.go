package main

import "fmt"

func main() {

	fmt.Println(foo())
	num, str := bar()
	fmt.Println(str, num)
}

func foo() int {
	return 65
}

func bar() (int, string) {
	return 65, "Andrews Football Number"
}
