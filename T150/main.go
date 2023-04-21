package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", server)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./web/css"))))
	http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("./web/pic"))))
	http.ListenAndServe(":8080", nil)
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./web/"))))
}

func server(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "web/")
}
