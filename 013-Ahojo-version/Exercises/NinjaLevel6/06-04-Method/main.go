package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %s %s, and I'm %d years old\n", p.first, p.last, p.age)
}

func main() {

	var kairi person = person{
		first: "Kairi",
		last:  "Hojnowski",
		age:   2,
	}

	kairi.speak()
}
