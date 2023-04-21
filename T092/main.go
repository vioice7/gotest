package main

import (
	"api"
	"net/http"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
