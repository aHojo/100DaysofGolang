package main

import "fmt"

func main() {
	sos := [][]string{
		{"James", "Bond", "Shaken not stirred"},
		{"Miss", "Moneypenney", "Hellooooooo, James"},
	}

	fmt.Println(sos)

	for _, v := range sos {
		for _, v2 := range v {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
}
