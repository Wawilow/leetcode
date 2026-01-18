package main

import "fmt"

func isMagic(grid [][]int) bool {
	eq := -1
	// rows
	row_sum := 0
	for _, row := range grid {
		row_sum = 0
		for _, v := range row {
			row_sum += v
		}
		if eq == -1 {
			eq = row_sum
		} else {
			if row_sum != eq {
				return false
			}
		}
	}

	// col
	col_sum := 0
	for i := 0; i < len(grid); i++ {
		col_sum = 0
		for _, row := range grid {
			col_sum += row[i]
		}

		if eq == -1 {
			eq = col_sum
		} else {
			if col_sum != eq {
				return false
			}
		}
	}

	if eq == -1 {
		return true
	}
	return true
}

func dfs(grid [][]int, size int) int {
	xm, ym := len(grid[0]), len(grid)

	if size <= 1 {
		return 1
	}

	for x := 0; x < xm-size+1; x++ {
		for y := 0; y < ym-size+1; y++ {
			gr := [][]int{}
			for _, v := range grid[y : y+size] {
				gr = append(gr, v[x:x+size])
			}
			if isMagic(gr) {
				fmt.Println(x, y, gr, size)
				return size
			}
		}
	}

	return dfs(grid, size-1)
}

func largestMagicSquare(grid [][]int) (res int) {
	res = 1
	larg := len(grid[0])
	if len(grid) > larg {
		larg = len(grid)
	}
	return dfs(grid, len(grid))
}
