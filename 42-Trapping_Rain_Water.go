package main

import "fmt"

type matrix_row struct {
	row []int
}

func maxHeight(height []int) int {
	ans := 0
	for _, v := range height {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func makeMatrix(height []int) ([]matrix_row, int) {
	res := []matrix_row{}

	for i := 0; i < maxHeight(height); i++ {
		row := matrix_row{}

		for j := 0; j < len(height); j++ {
			switch {
			case i < height[j]:
				row.row = append(row.row, 1)
			default:
				row.row = append(row.row, 0)
			}
		}

		res = append(res, row)
	}
	return res, len(height)
}

func contain1(list []int) bool {
	for _, b := range list {
		if b == 1 {
			return true
		}
	}
	return false
}

func waterMatrix(matrix []matrix_row, length int) []matrix_row {
	res := []matrix_row{}

	for i := 0; i < len(matrix); i++ {
		row := matrix_row{}

		for j := 0; j < len(matrix[i].row); j++ {
			if matrix[i].row[j] == 0 {
				left := false
				right := false

				left_list := []int{}
				right_list := []int{}
				if j != 0 {
					left_list = matrix[i].row[:(j)]
				}
				if j != length {
					right_list = matrix[i].row[(j):]
				}

				if contain1(left_list) == true {
					left = true
				}
				if contain1(right_list) == true {
					right = true
				}

				if left && right {
					row.row = append(row.row, 2)
				} else {
					row.row = append(row.row, 0)
				}
				//fmt.Println(left_list, right_list, left, right)
			} else {
				row.row = append(row.row, 0)
			}
		}

		res = append(res, row)
	}

	return res
}

func countWater(matrix []matrix_row) int {
	ans := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i].row); j++ {
			if matrix[i].row[j] == 2 {
				ans += 1
			}
		}
	}
	return ans
}

func trap(height []int) int {
	matrix, length := makeMatrix(height)
	waterM := waterMatrix(matrix, length)
	return countWater(waterM)
}

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
