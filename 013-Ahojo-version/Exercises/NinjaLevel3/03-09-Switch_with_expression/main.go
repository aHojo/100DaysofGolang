package main

import "fmt"

func main() {

	switch name := "Kairi"; name {
	case "Andrew":
		fmt.Println("Push yourself harder")
	case "Kairi":
		fmt.Println("You is perfect")
	default:
		fmt.Println("You must be lindsay")
	}
}
