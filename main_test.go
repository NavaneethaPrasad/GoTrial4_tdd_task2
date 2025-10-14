package main

import (
	"testing"
)

func TestLargest(t *testing.T) {
	actual := largestNum([]int{1, 2, 3})
	expected := 3
	if actual != expected {
		t.Errorf("Expected is %d but got %d", expected, actual)
	}
}

func TestLargest2(t *testing.T) {
	actual := largestNum([]int{9, 7, 4})
	expected := 9
	if actual != expected {
		t.Errorf("Expected is %d but got %d", expected, actual)
	}
}
