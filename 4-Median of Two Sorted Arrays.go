package main

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func bubbleSort(array []int) []int {
	var (
		result []int
	)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				//fmt.Print(j, " - ", array[j], array[j+1], "\n")
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	result = array
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		ans   float64
		slice = nums1
	)

	for _, v := range nums2 {
		slice = append(slice, v)
	}
	slice = bubbleSort(slice)

	if len(slice)%2 == 1 {
		ans = float64(slice[(len(slice) / 2)])
	} else {
		elem_one := float64(slice[len(slice)/2-1])
		elem_two := float64(slice[(len(slice) / 2)])
		ans = (elem_one + elem_two) / 2
	}

	return ans
}

//func main() {
//	fmt.Print(findMedianSortedArrays([]int{1}, []int{4}))
//}
