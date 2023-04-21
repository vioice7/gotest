package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Nume    string
	Adresa  string
	Oras    string
	Zip     string
	Regiune string
}

type hoteluri []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	roh := hoteluri{
		hotel{
			"Testo",
			"Str. Test, Nr. 4",
			"Bistrita",
			"420140",
			"Nord",
		},
		hotel{
			"Testo",
			"Str. Test, Nr. 40",
			"Cluj",
			"421140",
			"Cantru",
		},
		hotel{
			"Testo",
			"Str. Test, Nr. 124",
			"Targu Mures",
			"422140",
			"Sud",
		},
	}

	err := tpl.Execute(os.Stdout, roh)
	if err != nil {
		log.Fatalln(err)
	}
}
