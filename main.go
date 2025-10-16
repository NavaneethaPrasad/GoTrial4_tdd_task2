package main

import "fmt"

func analyzeNumbers(num []int) (int, int, float64) {
	largest := num[0]
	smallest := num[0]
	sum := 0

	for _, n := range num {
		if n > largest {
			largest = n
		}
		if n < smallest {
			smallest = n
		}
		sum += n
	}

	average := float64(sum) / float64(len(num))
	return largest, smallest, average
}

func main() {
	l, s, a := analyzeNumbers([]int{1, 2, 3})
	fmt.Println("Largest:", l)
	fmt.Println("Smallest:", s)
	fmt.Println("Average:", a)
}
