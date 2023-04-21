package main

import "fmt"

func main() {

	t0 := []string{"test0", "test1", "test2", "test3", "test4"}
	t1 := []string{"test5", "test6", "test7", "test8", "test9"}
	t2 := []string{"test10", "test11", "test12", "test13", "test14"}

	fmt.Println(t0)

	fmt.Println(t1)

	fmt.Println(t2)

	// multi dimensional slice
	t01 := [][]string{t0, t1, t2}

	fmt.Println(t01)
}
