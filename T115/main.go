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

func (p person) Test(x int) int {
	return x * 10
}

func (p person) SomeProcessing() int {
	return 10
}

func (p person) AgeDouble() int {
	return p.Age * 2
}

func (p person) ArgumentDouble(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p := person{
		Name: "James Bond",
		Age:  35,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}

// Generally speaking, best practice:
// call functions in templates for formatting only; not processing or data access.

// The main reasons you don't want to do any data processing in your template:
// (1) separation of concerns
// (2) if you're using a function more than once in a template,
// the server needs to do the processing more than once.
// (though the standard library might cache processing -
// I've yet to dig into this - if you find the answer, let me know)
