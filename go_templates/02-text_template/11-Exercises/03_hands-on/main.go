package main

import (
	"log"
	"os"
	"text/template"
)

type Hotel struct {
	Name, Address, City, Zip, Region string
}

type Region struct {
	Region string
	Hotels []Hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	regions := []Region{
		{
			Region: "Southern",
			Hotels: []Hotel{
				{
					Name:    "One",
					Address: "23423423 Maine street",
					City:    "Your mum",
					Zip:     "2323",
				},
				{
					Name:    "two",
					Address: "a223234 asdf street",
					City:    "Your dad",
					Zip:     "45454",
				},
				{
					Name:    "Three",
					Address: "232323 Maine street",
					City:    "Your Sister",
					Zip:     "23555523",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, regions)
	if err != nil {
		log.Fatalln(err)
	}

}
