package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := string(99)

	fmt.Println(s) // Prints out c because ascii 99 is c

	// s1 := string(44.2)
	var myStr = fmt.Sprintf("%f", 44.2) // Converts specifier to string
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 3434)
	fmt.Println(myStr1)
	fmt.Println(string(3434)) // ൪ prints out the unicode char

	// Convert String to numers need strconv package

	s1 := "3.234" // type string
	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64) // returns a float64
	_ = err
	fmt.Println(f1)
	fmt.Printf("%T\n", f1)

	//ascii to int;
	if i, err := strconv.Atoi("-50"); err == nil {
		fmt.Printf(" Int: %d, type = %T\n", i, i) //Int: -50, type = int
	}

	s2 := strconv.Itoa(55)
	fmt.Printf("String:%v, type: %T\n", s2, s2)
}
