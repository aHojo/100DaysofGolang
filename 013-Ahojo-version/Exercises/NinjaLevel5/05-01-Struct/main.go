package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first:   "Andrew",
		last:    "Hojnowski",
		flavors: []string{"Vanilla", "Cookie Dough", "Chocolate"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for _, v := range p1.flavors {
		fmt.Printf("%s ", v)
	}
	fmt.Println()
}
