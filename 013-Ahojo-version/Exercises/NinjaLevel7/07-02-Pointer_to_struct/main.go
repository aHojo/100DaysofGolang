package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var kairi *person = &person{
		name: "Kairi",
		age:  2,
	}
	fmt.Println("Struct before", kairi)
	changeMe(kairi)
	fmt.Println("Struct after", kairi)

}

func changeMe(p *person) {
	// (*struct).field to dereference a struct field
	(*p).name = "Kairi Hojnowski"
	// p.name is shorthand for this, the compiler can figure it out for us

}
