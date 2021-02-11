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
	sages := []Sage{
		{
			Name: "Andrew Hojo",
			Rank: "Most Powerful",
		},
		{
			Name: "Kairi Hojo",
			Rank: "2nd Powerful",
		},
		{
			Name: "Lindy",
			Rank: "Magic Only",
		},
		{
			Name: "Dog",
			Rank: "Good Boy",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {

		log.Fatalln(err)
	}
}
