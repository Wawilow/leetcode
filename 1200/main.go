package main

func bubbleSort(input []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] < input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func bubbleSort2(input [][]int) [][]int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			fmt.Println(i, input[i-1][0], input[i][0])
			if input[i-1][0] > input[i][0] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func abs(v int) int {
	if v >= 0 {
		return v
	} else {
		return v * -1
	}
}

func minimumAbsDifference(arr []int) [][]int {
	arr = bubbleSort(arr)
	current_min := -1
	res := [][]int{}
	for i, v := range arr {
		if i == len(arr)-1 {
			break
		}
		// fmt.Println(abs(arr[i+1]-v), current_min)
		if abs(arr[i+1]-v) < current_min || current_min == -1 {
			res = [][]int{[]int{arr[i+1], v}}
			current_min = abs(arr[i+1] - v)
		} else if abs(arr[i+1]-v) == current_min {
			res = append(res, []int{arr[i+1], v})
		}
	}
	return bubbleSort2(res)
}

func main() {}
