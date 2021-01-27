package main

import "fmt"

func main() {
	var years_old_count = 0
	start_year := 1990
	for {
		fmt.Println(start_year + years_old_count)
		years_old_count++
		if years_old_count >= 31 {
			break
		}
	}

}
