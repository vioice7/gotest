package main

import "fmt"

func main() {

	test := []string{"t0", "t1", "t2", "t3"}

	for _, v := range test {
		fmt.Println(v + "\n")

	}

	e, d := test01(9)

	fmt.Println(e, d)

}

func test01(a int) (int, int) {
	b := a - 10
	c := a + 10
	return b, c
}
