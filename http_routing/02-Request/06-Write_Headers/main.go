package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("HojoKey", "I'm the GOAT")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintln(w, "<h1>Can add the html here</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8000", d)
}
