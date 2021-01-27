package main

import "fmt"

func main() {
	for i := 4; i <= 100; i++ {
		if i%4 == 0 {

			fmt.Println(i)
		}
	}
}
