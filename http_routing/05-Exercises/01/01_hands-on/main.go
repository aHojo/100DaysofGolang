package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "MAIN INDEX PAGE %s", "HI")

}
func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>I am dog</h1>")
}
func cat(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>I am cat</h1>")
}
func me(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>Andrew Hojnowski</h1>")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)
	http.HandleFunc("/me/", me)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
