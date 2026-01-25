package main

import "fmt"

func bubbleSort(input []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func minimumDifference(nums []int, k int) (res int) {
	if k == 1 {
		return
	}

	// thought 1
	// bubble sort the array -> go through it with the k spacing between elems?

	nums = bubbleSort(nums)
	// fmt.Println(nums)
	for i := 0; i <= len(nums)-k; i++ {
		// fmt.Println(nums[i+k-1] - nums[i], nums[i+k-1], nums[i], i, k)
		if res == 0 || nums[i+k-1]-nums[i] < res {
			res = nums[i+k-1] - nums[i]
		}
	}
	return
}

func main() {
	fmt.Println([]int{9, 4, 1, 7}, 2)
}
