package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	// using the default server multiplexor (nil to ListenAndServe)
	http.Handle("/dog/", d)
	http.Handle("/cat/", c)

	http.ListenAndServe(":8080", nil)
}
