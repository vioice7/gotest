package main

import (
	"io"
	"net/http"
)

func id(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is the index.")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is the dog module.")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "My name is: Vio")
}

func main() {

	// using the default server multiplexor (nil to ListenAndServe)
	http.HandleFunc("/", id)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
