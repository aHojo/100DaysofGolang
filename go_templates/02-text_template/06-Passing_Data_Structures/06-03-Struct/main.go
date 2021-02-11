package main

import (
	"log"
	"os"
	"text/template"
)

type Sage struct {
	Name string
	Rank string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sage := Sage{
		Name: "Andrew Hojo",
		Rank: "Most Powerful",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sage)
	if err != nil {

		log.Fatalln(err)
	}
}
