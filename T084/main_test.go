package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected 5 got", x)
	}
}

func TestMySubstract(t *testing.T) {
	x := mySubstract(10, 3)
	if x != 7 {
		t.Error("Expected 7 got", x)
	}
}
