package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/dog":
		fmt.Fprintln(w, "<h1>Dogs are the best</h1>")
	case "/cat":
		fmt.Fprintln(w, "<h2>Ew a cat</h2>")
	default:
		fmt.Fprintln(w, `
<ul>
	<li><a href="/dog">Dog</li>
	<li><a href="/cat">Cat</li>
</ul>
		`)
	}
}
func main() {
	var d hotdog

	http.ListenAndServe(":8000", d)

}
