package main

import "fmt"

func main() {
	var f func() = func() {
		fmt.Println("I'm in the function expression")
	}

	f()
}
