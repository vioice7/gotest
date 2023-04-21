package main

import (
	"fmt"
	"log"
	"net"
)

// utilizeaza celalat server pentru trimitere informatii (scriere)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "Te am sunat pe tine.")
}
