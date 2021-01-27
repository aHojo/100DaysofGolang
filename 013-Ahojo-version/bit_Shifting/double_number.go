package main

import (
	"fmt"
)

func main() {
	var x = 2
	fmt.Printf("Decimal > %d\t\tBinary > %b\n", x, x)
	
	// This shift to the left one, basically just multiplies by 2
	var y = x << 1 // shift the bit to the left 1
	fmt.Printf("Decimal > %d\t\tBinary > %b\n", y, y)
	var z = y << 1 // shift the bit to the left 1 
	fmt.Printf("Decimal > %d\t\tBinary > %b\n", z, z)
}

