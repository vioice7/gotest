package main

import "fmt"

func test01() {
	fmt.Println("t01")
}

func main() {

	test01()
	//

	// anonymoous function
	func() {
		fmt.Println("Anonymous function ran.")
	}()
	func(x int) {
		fmt.Println("Anonymous function ran. ", x)
	}(12)

}
