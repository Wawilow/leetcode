package main

import (
	"fmt"
	"strconv"
)

func largestNumber(nums []int) (r string) {
	allO := true
	for _, n := range nums {
		if n != 0 {
			allO = false
			break
		}
	}
	if allO {
		return "0"
	}

	ll := make([]string, len(nums))
	for i, v := range nums {
		ns := strconv.Itoa(v)
		ll[i] = ns
	}

	for i := 0; i < len(ll)-1; i++ {
		for j := 0; j < len(ll)-i-1; j++ {
			if !compare(ll[j], ll[j+1]) {
				ll[j], ll[j+1] = ll[j+1], ll[j]
			}
		}
	}

	for _, v := range ll {
		r += v
	}

	return
}

func compare(s1, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 == l2 {
		return s1 > s2
	}
	return s1+s2 > s2+s1
}

func main() {
	//fmt.Println(largestNumber([]int{39, 10, 2}))
	fmt.Println(largestNumber([]int{0, 0, 0, 0}))
}
