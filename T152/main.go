package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("public/templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", test)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public/pics"))))
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
