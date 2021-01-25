package main

import "fmt"

type account int

var x account
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 1000000
	fmt.Println(x)

	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
