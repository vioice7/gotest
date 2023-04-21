package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type cars struct {
	Model     string
	Type      string
	Automatic bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	c1 := cars{
		Model:     "Audi",
		Type:      "A8",
		Automatic: true,
	}

	c2 := cars{
		Model:     "BMW",
		Type:      "530d",
		Automatic: true,
	}

	c3 := cars{
		Model:     "Hyundai",
		Type:      "i40",
		Automatic: false,
	}

	c4 := cars{
		Model:     "Honda",
		Type:      "Accord",
		Automatic: false,
	}

	users := []cars{c1, c2, c3, c4}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
