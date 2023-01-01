package _

func twoSum(nums []int, target int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i != j {
				if nums[i]+nums[j] == target {
					ans = []int{i, j}
				}
			}
		}
	}
	return ans
}

//func main() {
//	fmt.Print("Start\n")
//	nums := []int{3, 2, 4}
//	target := 6
//	r := twoSum(nums, target)
//	fmt.Print(r, "\n")
//	fmt.Print("End\n")
//}
