package main

import "fmt"

func main() {

	c := make(chan int)

	go foo(c)

	bar(c)

	fmt.Println("Exiting....")
}

func foo(c chan<- int) {
	defer close(c)
	for i := 0; i < 100; i++ {

		c <- i
	}

}

func bar(c <-chan int) {
	for v := range c {

		fmt.Println(v)
	}
}
