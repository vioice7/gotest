package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type house struct {
	Rooms     int
	Bathrooms int
	Type      string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	b := house{
		Rooms:     4,
		Bathrooms: 2,
		Type:      "Apartment",
	}

	g := house{
		Rooms:     2,
		Bathrooms: 1,
		Type:      "Apartment",
	}

	m := house{
		Rooms:     10,
		Bathrooms: 4,
		Type:      "Luxury Home",
	}

	f := car{
		Manufacturer: "Hyundai",
		Model:        "Tucson",
		Doors:        4,
	}

	c := car{
		Manufacturer: "Mitsubishi",
		Model:        "Lancer",
		Doors:        4,
	}

	houses := []house{b, g, m}
	cars := []car{f, c}

	data := struct {
		Houses    []house
		Transport []car
	}{
		houses,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
