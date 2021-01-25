package main

import "fmt"

type partners int

var x partners

func main() {
	fmt.Printf("%d\nType: %T\n", x, x)
	x = 10005
	fmt.Println(x)

}
