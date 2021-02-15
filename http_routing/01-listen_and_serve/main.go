package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Any code you want in this func")
}
func main() {

	d := 42
	err := http.ListenAndServe("localhost:8000", hotdog(d))
	if err != nil {
		log.Fatalln(err)
	}

}
