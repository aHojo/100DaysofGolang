package main

import "fmt"

func main() {
	name := "Kairi"

	switch {
	case name == "Andrew":
		fmt.Println("Push yourself harder")
	case name == "Kairi":
		fmt.Println("You is perfect")
	default:
		fmt.Println("You must be lindsay")
	}
}
