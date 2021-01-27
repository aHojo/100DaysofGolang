package main

import "fmt"

func main() {
	start := 65
	stop := 90

	for i := start; i <= stop; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

}
