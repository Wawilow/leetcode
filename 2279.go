package main

import "fmt"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	res := 0
	bfrFillIn := make([]int, len(capacity))

	for i, _ := range capacity {
		bfrFillIn[i] = capacity[i] - rocks[i]
	}
	for i := 0; i < len(bfrFillIn)-1; i++ {
		for j := 0; j < len(bfrFillIn)-i-1; j++ {
			if bfrFillIn[j] > bfrFillIn[j+1] {
				bfrFillIn[j], bfrFillIn[j+1] = bfrFillIn[j+1], bfrFillIn[j]
			}
		}
	}
	for additionalRocks >= 0 {
		//fmt.Println(additionalRocks)
		if len(bfrFillIn) == 0 {
			return res
		} else if bfrFillIn[0] <= additionalRocks {
			additionalRocks -= bfrFillIn[0]
			bfrFillIn = bfrFillIn[1:]
			res += 1
		} else {
			return res
		}
	}
	//fmt.Println("bfr fill in:", bfrFillIn)
	return res
}

func main() {
	//fmt.Println("ans: ", maximumBags([]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2))
	fmt.Println("ans: ", maximumBags([]int{}, []int{}, 100))
}
