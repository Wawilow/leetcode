package main

import (
	"fmt"
)

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

func dfs(a []int, v, sum int) int {
	if sum == v {
		return sum
	}
	if len(a) == 0 {
		return -1
	}
	if sum+a[0] == v {
		return sum + a[0]
	}
	if a[0] > v || sum+a[0] > v {
		return -1
	}

	// don't include current val => smaller overall result
	res := dfs(a[1:], v, sum)
	if res != -1 {
		return res
	}
	return dfs(a[1:], v, sum+a[0])
}

func anotherDfs(l1, l2 []int, diff int) int {
	if diff < 0 {
		l1, l2 = l2, l1
	}

	res := dfs(l1, diff, 0)
	if res != -1 {
		return res
	}

	res = anotherDfs(l1, l2[1:], diff)

	fmt.Println(diff, l1, l2)
	return -1
}

func minimumDeleteSum(s1 string, s2 string) int {
	l1 := []int{}
	l2 := []int{}
	var sum1, sum2 int
	for _, v := range s1 {
		l1 = append(l1, int(v))
		sum1 += int(v)
	}
	for _, v := range s2 {
		l2 = append(l2, int(v))
		sum2 += int(v)
	}
	if sum1 == sum2 {
		return 0
	}
	l1 = bubbleSort(l1)
	l2 = bubbleSort(l2)

	if sum1-sum2 < 0 {
		l1, l2 = l2, l1
	}

	r1 := anotherDfs(l1, l2, sum1-sum2)
	if r1 == -1 {
		return sum1 + sum2
	}
	return r1
}

func main() {
	fmt.Println("final: ", minimumDeleteSum("delete", "leet"))
	// fmt.Println("final: ", minimumDeleteSum("eat", "sea"))
}
