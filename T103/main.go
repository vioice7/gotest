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

	cars := []string{"Audi", "BMW", "Mitsubishi", "Hyundai", "Honda", "Volkswagen"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", cars)
	if err != nil {
		log.Fatalln(err)
	}

}
