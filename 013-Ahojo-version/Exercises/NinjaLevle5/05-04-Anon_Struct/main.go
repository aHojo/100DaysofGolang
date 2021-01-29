package main

import "fmt"

func main() {
	kairi := struct {
		name string
		last string
		age  int
	}{
		name: "Kairi",
		last: "Hojnowski",
		age:  2,
	}

	fmt.Println(kairi)
}
