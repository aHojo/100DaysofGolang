package main

import "fmt"

type human interface {
	speak()
}
type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Printf("My name is %s, and I am %d years old\n", p.name, p.age)
}

func saySomething(h human) {
	h.speak()
}
func main() {

	var p1 *person = &person{
		name: "kairi",
		age:  2,
	}

	p2 := person{
		name: "andrew",
		age:  30,
	}

	saySomething(p1)

	//  below should not work.
	//saySomething(p2)
	/*
		Error Message:
			cannot use p2 (variable of type person) as human value in argument to saySomething: missing method speak (speak has pointer receiver)compiler
	*/
	// CAN do this
	p2.speak()
}
