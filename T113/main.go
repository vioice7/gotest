package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type agent struct {
	person
	LicenceToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	a1 := agent{
		person{
			Name: "James Bond",
			Age:  42,
		},
		true,
	}

	err := tpl.Execute(os.Stdout, a1)
	if err != nil {
		log.Fatalln(err)
	}

}
