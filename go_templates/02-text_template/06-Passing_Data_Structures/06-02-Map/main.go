package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := map[string]string{
		"Most Powerful":        "Andrew Hojo",
		"Second Most Powerful": "Kairi Hojo",
		"Only has magic":       "Lindy",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {

		log.Fatalln(err)
	}
}
