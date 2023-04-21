package main

import (
	"io"
	"net/http"
)

func hotdog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func hotcat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	// using the default server multiplexor (nil to ListenAndServe)
	http.HandleFunc("/dog/", hotdog)
	http.HandleFunc("/cat/", hotcat)

	http.ListenAndServe(":8080", nil)
}
