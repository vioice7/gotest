package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

var test = template.FuncMap{
	"uc": strings.Title,
	"ft": firstThree,
	"lt": lastThree,
	"rv": sReverse,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func lastThree(s string) string {
	s = strings.TrimSpace(s)
	s = sReverse(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	s = sReverse(s)
	return s
}

func sReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func init() {
	tpl = template.Must(template.New("").Funcs(test).ParseFiles("tpl.gohtml"))
}

func main() {

	a := car{
		Manufacturer: "hyundai",
		Model:        "tucson",
		Doors:        4,
	}

	b := car{
		Manufacturer: "mitsubishi",
		Model:        "lancer",
		Doors:        4,
	}

	c := car{
		Manufacturer: "honda",
		Model:        "accord",
		Doors:        4,
	}

	cars := []car{a, b, c}

	data := struct {
		Transport []car
	}{
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
