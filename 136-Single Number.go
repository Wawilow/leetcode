package main

func singleNumber(nums []int) int {
	var dict = make(map[int]int)

	for _, v := range nums {
		dict[v] = dict[v] + 1
	}
	for n, v := range dict {
		if v == 1 {
			return n
		}
	}
	return 0
}

func main() {
	singleNumber([]int{1, 2, 2})
}
