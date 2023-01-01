package _67

//func alrIn(vrb int, list []int) bool {
//	for _, v := range list {
//		if v == vrb {
//			return true
//		}
//	}
//	return false
//}

//func numsSameConsecDiff(n int, k int) []int {
//	res := []int{}
//
//	for i := 0; i < 10; i++ {
//		dfs(&res, []int{i}, n, k)
//	}
//
//	return res
//}

//func dfs(res *[]int, number []int, n int, k int) {
//	if len(number) == n {
//		// if num is already normal length append to result and end function
//		if number[0] != 0 {
//			// if first num is not zero
//			if !(alrIn(make_num(number), *res)) {
//				// if num already not in result list
//				*res = append(*res, make_num(number))
//			}
//		}
//		return
//	}
//
//	if 0 <= number[(len(number)-1)]-k && number[(len(number)-1)]-k < 10 {
//		num := number
//		num = append(num, (number[(len(number)-1)] - k))
//		dfs(res, num, n, k)
//	}
//
//	if 0 <= number[(len(number)-1)]+k && number[(len(number)-1)]+k < 10 {
//		num := number
//		num = append(num, (number[(len(number)-1)] + k))
//		dfs(res, num, n, k)
//	}
//
//	return
//}

//func make_num(list []int) int {
//	num := 0
//
//	//for i, v := range list {
//	for i := len(list) - 1; i >= 0; i-- {
//		v := list[len(list)-i-1]
//		num += v * int(math.Pow(10, float64(i)))
//
//	}
//	return num
//}

//func main() {
//	fmt.Println(time.Now())
//	fmt.Println(numsSameConsecDiff(2, 0))
//fmt.Println(make_num([]int{2, 2, 8}))
//fmt.Println(alrIn(3, []int{1, 0, 2}))
//fmt.Println(time.Now())
//}
