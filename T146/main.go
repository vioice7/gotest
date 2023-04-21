package main

import (
	"io"
	"net/http"
)

func main() {
	handler := http.FileServer(http.Dir("."))
	http.Handle("/", handler)
	http.HandleFunc("/car/", car)
	http.ListenAndServe(":8080", nil)
}

func car(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/car.jpg">`)
}
