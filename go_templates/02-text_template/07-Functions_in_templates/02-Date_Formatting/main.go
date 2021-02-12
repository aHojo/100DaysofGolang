package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmpl *template.Template

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func init() {
	tmpl = template.Must(template.New("doit").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {

	err := tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
