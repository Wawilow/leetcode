package main

import (
	"fmt"
)

func bubbleSort(input []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] < input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func maxSumDivThree(nums []int) (res int) {
	sorted := map[int][]int{
		1: {},
		2: {},
	}
	for _, v := range nums {
		switch v % 3 {
		case 0:
			res += v
		case 1:
			sorted[1] = append(sorted[1], v)
		case 2:
			sorted[2] = append(sorted[2], v)
		}
	}

	larger := 0
	if len(sorted[1]) >= len(sorted[2]) {
		larger = 1
	} else {
		larger = 2
	}
	sorted[larger] = bubbleSort(sorted[larger])

	for i, v := range sorted[3-larger] {
		res += v + sorted[larger][i]
	}
	return
}

func main() {
	fmt.Println("hi")
}
