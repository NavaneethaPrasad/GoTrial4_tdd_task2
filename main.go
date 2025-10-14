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

func main() {
	fmt.Println(largestNum([]int{1, 2, 3}))

}
