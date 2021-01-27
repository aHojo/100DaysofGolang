package main


import (
	"fmt"
)

func main(){
	var a string = `hello
	this string is a 
	raw string
	literal and preserves the
	spaces
	`
	fmt.Printf("%s\n", a)
}
