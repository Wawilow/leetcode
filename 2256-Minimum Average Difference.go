package main

func average(nums []int) int {
	length := len(nums)
	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	if length == 0 {
		return 0
	} else {
		return sum / length
	}
}

func averageDifference(nums1 []int, nums2 []int) int {
	res := (average(nums1) - average(nums2))
	if res < 0 {
		return res * -1
	} else {
		return res
	}
}

func minimumAverageDifference(nums []int) int {
	n := len(nums)

	res := 9999999999
	resIndex := 0

	for i := 0; i < n; i++ {
		//fmt.Println(i+1, " - ", n-i-1)
		//fmt.Println("! ", nums[:i+1], " - ", nums[n-(n-i-1):])
		aD := averageDifference(nums[:i+1], nums[n-(n-i-1):])
		//fmt.Println("aD -", aD, nums[:i+1], nums[n-(n-i-1):])
		if aD < res {
			resIndex = i
			res = aD
		}
	}
	return resIndex
}

//func main() {
//	fmt.Println(minimumAverageDifference([]int{2, 5, 3, 9, 5, 3}))
//	fmt.Print("that's end")
//}
