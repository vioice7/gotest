package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func id(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is the index.")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is the dog module.")
}

func me(res http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "tpl.gohtml", "McLeod")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {

	// using the default server multiplexor (nil to ListenAndServe)
	http.HandleFunc("/", id)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
