package main

import "fmt"

func findLHS(nums []int) int {
	mapping := map[int]int{}
	for _, v := range nums {
		mapping[v] += 1
	}

	max := 0
	for i := len(mapping); i > 0; i-- {
		booba, boba := mapping[i], mapping[i-1]
		if booba == 0 || boba == 0 {
			continue
		}
		if booba+boba > max {
			max = boba + booba
		}
	}
	return max
}

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	// fmt.Println(findLHS([]int{1, 2, 3, 4, 6}))
}
