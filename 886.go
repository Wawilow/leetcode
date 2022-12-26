package main

//func possibleBipartition(n int, dislikes [][]int) bool {
//	var (
//		colors = make(map[int]int)
//		nums   = make(map[int][]int)
//	)
//	for i := 1; i <= n; i++ {
//		nums[i] = append(nums[i])
//		colors[i] = 0
//	}
//	for _, v := range dislikes {
//		nums[v[1]] = append(nums[v[1]], v[0])
//		nums[v[0]] = append(nums[v[0]], v[1])
//	}
//	for i := 1; i <= n; i++ {
//		if colors[i] != 0 {
//			continue
//		} else {
//			if !dfs(i, 1, colors, nums) {
//				return false
//			}
//		}
//	}
//	return true
//}

//func dfs(n int, group int, colors map[int]int, nums map[int][]int) bool {
//	fmt.Println(n, group)
//	if colors[n] == group {
//		return true
//	}
//	if colors[n] == 3-group {
//		return false
//	}
//	if colors[n] == 0 {
//		colors[n] = group
//		for _, v := range nums[n] {
//			if !dfs(v, 3-group, colors, nums) {
//				return false
//			}
//		}
//	}
//	return true
//}

//func main() {
//	print(possibleBipartition(4, [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 4}}))
//}
