package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)

}

func sroot(w http.ResponseWriter, r *http.Request) {

	culori := []string{"00AA00", "AA0000", "0000AA", "AAAA00", "AA00AA", "00AAAA"}

	titlu := "Test"
	cap := "<html><head><title>" + titlu + "</title></head><body>"
	coada := "</body></html>"
	corp := ""

	for i := 0; i < 1000; i++ {
		if i%1 == 0 {
			corp += " <strong style=\"color:" + culori[rand.Intn(6)] + ";\">test</strong> "
		} else {
			corp += " test "
		}

	}

	pag := cap + corp + coada

	fmt.Fprintf(w, pag)
}
