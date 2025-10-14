package main

import (
	"fmt"
	"testing"
)

func TestLargest(t *testing.T) {
	actual := largest([]int{1, 2, 3})
	expected := 3
	if actual != expected {
		fmt.Errorf("Expected is %d but got %d", expected, actual)
	}
}
