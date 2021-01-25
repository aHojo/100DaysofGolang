package main

import "fmt"

var x uint16 = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%d\t%s\t%t\n", x, y, z)

	fmt.Println(s)

}
