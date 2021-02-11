package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Sage struct {
	Name string
	Rank string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
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
