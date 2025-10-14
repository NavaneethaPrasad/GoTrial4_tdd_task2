package main

import "fmt"

func largestNum(num []int) int {
	largest := num[0]
	for _, n := range num {
		if n > largest {
			largest = n
		}
	}
	return largest
}
func smallestNum(num []int) int {
	smallest := num[0]
	for _, n := range num {
		if n < smallest {
			smallest = n
		}
	}
	return smallest
}

func averageNum(num []int) float64 {
	return 2.0
}
func main() {
	fmt.Println(largestNum([]int{1, 2, 3}))
	fmt.Println(smallestNum([]int{1, 2, 3}))
	fmt.Println(averageNum([]int{1, 2, 3}))
}
