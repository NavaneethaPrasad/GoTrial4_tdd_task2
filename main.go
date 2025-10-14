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
	sum := 0
	for _, n := range num {
		sum += n
	}
	average := float64(sum) / float64(len(num))
	return average
}
func main() {
	fmt.Println(largestNum([]int{1, 2, 3}))
	fmt.Println(smallestNum([]int{1, 2, 3}))
	fmt.Println(averageNum([]int{1, 2, 3}))
}
