package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {

	p1 := person{
		First: "Kairi",
		Last:  "Hojnowski",
		Sayings: []string{
			"Kairi",
			"Maow",
			"QUACK!!!!",
		},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(bs))
}
