package main

import "fmt"

func main() {

	c := make(chan int, 1) // buffered channel - 1 value is allowed to sit in here

	c <- 42

	fmt.Println(<-c)
}
