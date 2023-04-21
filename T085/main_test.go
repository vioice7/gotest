package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	// a test is going to be a slice of type test
	// a slice of values of type test
	tests := []test{
		// composit literal (a slice of int) and an int
		test{[]int{5, 9}, 14},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		// unfurl the slice of ints
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}

}
