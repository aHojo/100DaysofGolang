package main

import "fmt"

func main() {

	defer foo()
	bar()
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}
