package main

import "tawesoft.co.uk/go/dialog"

func main() {
	for i := 0; i < 10; i++ {
		dialog.Alert("Test %d", i)
	}

}
