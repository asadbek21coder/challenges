package main

import (
	"fmt"
	// "math"
)

func main() {
	fmt.Println(MinArrayCount([]int{8,8,8, 2, 2, 6, 5, 3, 2, 5, 8, 9, 10, 11, 2, 2, 2, 3, 2, 2, 4, 2, 2, 5, 6, 7, 8, 9}))
}

func MinArrayCount(array []int) (min, count int) {
	min = array[0]
	for _, v := range array {
		if v < min {
			min = v
			count = 0
		}
		if v == min {
			count++
		}
	}
	return min, count
}
