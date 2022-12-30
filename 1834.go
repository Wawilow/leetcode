package main

import (
	"fmt"
)

func getOrder(tasks [][]int) []int {
	res := []int{}
	avaiable := []int{}
	cur := 0
	for i := 0; i <= tasks[0][0]; i++ {
		//fmt.Println(cur)
		if cur+i < len(tasks) {
			avaiable = append(avaiable, cur+i)
		}
	}

	dfs(&avaiable, &tasks, &res)

	return res
}
func dfs(avaiable *[]int, tasks *[][]int, res *[]int) {
	if len(*tasks) == 0 {
		return
	}

	if len(*res) == len(*tasks) {
		return
	} else {
		for i := 0; i < len(*avaiable)-1; i++ {
			for j := 0; j < len(*avaiable)-i-1; j++ {
				//fmt.Println("j: ", j, (*avaiable)[j], (*avaiable)[j+1])
				if (*tasks)[(*avaiable)[j]][1] > (*tasks)[(*avaiable)[j+1]][1] {
					(*avaiable)[j], (*avaiable)[j+1] = (*avaiable)[j+1], (*avaiable)[j]
				}
			}
		}
		// now continue
		fmt.Println(res, avaiable)
		cur := maxOne(*res, *avaiable)
		solv := (*avaiable)[0]
		*res = append(*res, (*avaiable)[0])
		*avaiable = (*avaiable)[1:]
		fmt.Println(res, avaiable)
		fmt.Println()

		//fmt.Println("111: ", (*tasks)[solv])
		for i := 0; i <= ((*tasks)[solv][0])+0; i++ {
			//fmt.Println(cur)
			if cur+i < len(*tasks) {
				*avaiable = append(*avaiable, cur+i)
			}
		}
		dfs(avaiable, tasks, res)
	}
}

func maxOne(ready []int, avaiable []int) int {
	res := 0
	for _, v := range ready {
		if v > res {
			res = v
		}
	}
	for _, v := range avaiable {
		if v > res {
			res = v
		}
	}
	return res
}

//func delete(i, a []int) []int {
//	return a[:i+copy(a[i:], a[i+1:])]
//}

func main() {
	fmt.Println(getOrder([][]int{[]int{7, 10}, []int{7, 12}, []int{7, 5}, []int{7, 4}, []int{7, 2}}))
	//fmt.Println(len([]int{0, 1, 2, 3}), []int{2, 4})
}
