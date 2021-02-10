package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")

	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, nil) // io.Writer, and a data

	if err != nil {
		log.Fatal(err)
	}

}
