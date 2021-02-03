package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go func() {
		fmt.Println("I'm printing something here 1 ")
		wg.Done()
	}()
	go func() {
		fmt.Println("I'm printing something here 2 ")
		wg.Done()
	}()
	fmt.Println("Number of goroutines running", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Exiting")
}

/*
Output:
	Number of goroutines running 3
	I'm printing something here 2
	I'm printing something here 1
	Exiting
*/
