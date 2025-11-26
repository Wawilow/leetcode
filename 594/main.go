package main

import "fmt"

func findLHS(nums []int) int {
	mapping := map[int]int{}
	for _, v := range nums {
		mapping[v] += 1
	}

	fmt.Println(len(mapping), mapping)

	max := 0
	for i := len(mapping); i > 0; i-- {
		booba, boba := mapping[i], mapping[i-1]
		fmt.Println(booba, boba)
		if booba == 0 || boba == 0 {
			continue
		}
		// booba *= i
		// boba *= i - 1
		if booba+boba > max {
			max = boba + booba
		}
	}
	// for key, val := range mapping {
	// 	if val+mapping[key+1] > max {
	// 		max = val + mapping[key+1]
	// 	}
	// 	if key != 0 {
	// 		if val+mapping[key-1] > max {
	// 			max = val + mapping[key-1]
	// 		}
	// 	}
	// }

	return max
}

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	// fmt.Println(findLHS([]int{1, 2, 3, 4, 6}))
}
