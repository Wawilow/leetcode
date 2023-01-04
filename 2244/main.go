package main

import "fmt"

func minimumRounds(tasks []int) int {
	res := -1

	count := 0
	dict := map[int]int{}
	for _, v := range tasks {
		dict[v] += 1
	}
	for i, v := range dict {
		fmt.Println(i, ":  ", v)
		if v%3 == 0 {
			count += v / 3
		} else if v%2 == 0 && v <= 4 {
			count += v / 2
		} else {
			num := v
			for num >= 2 {
				if num%2 == 0 && num <= 4 {
					num -= 2
					count += 1
				} else {
					num -= 3
					count += 1
				}
			}
			if num != 0 {
				return res
			}
		}
		//fmt.Println("count: ", count)
	}

	if count != 0 {
		res = count
	}
	return res
}

func main() {
	//fmt.Println(minimumRounds([]int{2, 2}))
	fmt.Println(minimumRounds([]int{69, 65, 62, 64, 70, 68, 69, 67, 60, 65, 69, 62, 65, 65, 61, 66, 68, 61, 65, 63, 60, 66, 68, 66, 67, 65, 63, 65, 70, 69, 70, 62, 68, 70, 60, 68, 65, 61, 64, 65, 63, 62, 62, 62, 67, 62, 62, 61, 66, 69}))
}
