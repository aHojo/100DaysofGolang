package main

import "fmt"

func main() {
	fmt.Println("Hello there Andrew")

	foo()

	for i := 0; i < 50; i++ {
		fmt.Printf("The var i is value %d\n", i)

		if i%2 == 0 {
			fmt.Printf("It's also even \n")
		}
	}

	bar()
	return
}

func foo() {
	fmt.Println("I'm in the function foooooooooo!!!1")
}

func bar() {
	fmt.Println("We exited")

}
