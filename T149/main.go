package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.HandleFunc("/car/", car)
	http.HandleFunc("/car.jpg", carimg)
	http.ListenAndServe(":8080", nil)
}

func car(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func carimg(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "car.jpg")
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `foo ran`)
}
