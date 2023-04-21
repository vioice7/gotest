package main

import (
	"io"
	"net/http"
)

const form = `<html><body>
<form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form>
</body></html>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>Hello!</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		content := form + request.FormValue("in")
		io.WriteString(w, content)
	}
}

func main() {
	http.HandleFunc("/test0", SimpleServer)
	http.HandleFunc("/test1", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
