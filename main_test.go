package main

import (
	"math"
	"testing"
)

func TestLargest(t *testing.T) {
	l, _, _ := analyzeNumbers([]int{1, 2, 3})
	expected := 3
	if l != expected {
		t.Errorf("Expected is %d but got %d", expected, l)
	}
}

func TestLargest2(t *testing.T) {
	l, _, _ := analyzeNumbers([]int{9, 7, 4})
	expected := 9
	if l != expected {
		t.Errorf("Expected is %d but got %d", expected, l)
	}
}

func TestSmallest(t *testing.T) {
	_, s, _ := analyzeNumbers([]int{1, 2, 3})
	expected := 1
	if s != expected {
		t.Errorf("Expected is %d but got %d", expected, s)
	}
}

func TestSmallest2(t *testing.T) {
	_, s, _ := analyzeNumbers([]int{4, 5, 2})
	expected := 2
	if s != expected {
		t.Errorf("Expected is %d but got %d", expected, s)
	}
}

func TestAverage(t *testing.T) {
	_, _, a := analyzeNumbers([]int{1, 2, 3})
	expected := 2.0
	if a != expected {
		t.Errorf("Expected is %.2f but got %.2f", expected, a)
	}
}

func TestAverage2(t *testing.T) {
	_, _, a := analyzeNumbers([]int{5, 8, 6})
	expected := 6.33
	a = math.Round(a*100) / 100
	if a != expected {
		t.Errorf("Expected is %.2f but got %.2f", expected, a)
	}
}
