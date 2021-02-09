package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive(c, q <-chan int) {

	for {

		select {
		case msg1 := <-c:
			fmt.Println(msg1)
		case msg2 := <-q:
			fmt.Println(msg2)
			return

		}
	}
}
