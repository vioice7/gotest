package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	nume := "VioMI"

	tpl := `<!DOCTYPE html>
	<html>
	<head>
	<title>` + nume + `</title>
	</head>
	<body>
	
	<h1>Salut ` + nume + `</h1>
	
	</body>
	</html>`

	fisier, err := os.Create("index.html")
	if err != nil {
		log.Fatal("eroare la creere fisier")
	}
	defer fisier.Close()

	io.Copy(fisier, strings.NewReader(tpl))
}
