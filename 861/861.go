package main

import (
	"fmt"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func matrixCount(grid [][]int) int64 {
	num := int64(0)
	for _, v := range grid {
		s := ""
		for _, sv := range v {
			s += strconv.Itoa(sv)
		}
		vNum, _ := strconv.ParseInt(s, 2, 64)
		num += vNum
	}
	return num
}

func matrixInvertRow(row []int) []int {
	res := []int{}
	for _, v := range row {
		res = append(res, Abs(v-1))
	}
	return res
}

func matrixScore(grid [][]int) int {
	xLen := len(grid)
	yLen := 0
	for i, row := range grid {
		if yLen == 0 {
			yLen = len(row)
		}
		invert := matrixInvertRow(row)
		if matrixCount([][]int{row}) < matrixCount([][]int{invert}) {
			grid[i] = invert
		}
	}

	for i := 0; i < yLen; i++ {
		sum := 0
		for _, row := range grid {
			sum += row[i]
		}

		if sum < xLen-sum {
			for _, row := range grid {
				row[i] = Abs(row[i] - 1)
			}
		}
	}
	return int(matrixCount(grid))
}

func main() {
	fmt.Println(matrixScore([][]int{
		[]int{0, 0, 1, 1},
		[]int{1, 0, 1, 0},
		[]int{1, 1, 0, 0},
	}))

	//matrixScore([][]int{
	//	[]int{1, 1, 1, 1},
	//	[]int{0, 0, 0, 1},
	//	[]int{1, 1, 1, 1},
	//})
}
