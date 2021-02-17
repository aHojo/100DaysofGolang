package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>Dog's are man's best friend</h1>")
}
func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h2>Ew it's a cat</h2>")
}
func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintln(w, `
		<ul>
			<li><a href="/dog">Dog</li>
			<li><a href="/cat">Cat</li>
		</ul>
				`)
	})
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8000", mux)

}
