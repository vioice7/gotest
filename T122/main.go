package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

// utilizeaza celalat server pentru preluare informatii (citire)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
