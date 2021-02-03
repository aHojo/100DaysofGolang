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

	m := map[string]person{
		p1.last: p1,
	}
	fmt.Println(m["Hojnowski"].first)
	fmt.Println(m["Hojnowski"].last)
	for _, v := range m["Hojnowski"].flavors {
		fmt.Printf("%s ", v)
	}
	fmt.Println()
}
