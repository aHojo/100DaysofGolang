package main

import "fmt"

const x string = "Your mom"

const y = 44

func main(){
	fmt.Printf("%s\n",x)
	fmt.Printf("%T\n",x)
	fmt.Printf("%d\n",y)
	fmt.Printf("%T\n",y)
}
