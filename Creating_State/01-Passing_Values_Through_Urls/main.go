package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	v := req.FormValue("q")
	io.WriteString(w, "Do my search Query: "+v)
}
