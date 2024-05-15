package main

import "fmt"

func dfs(mg *int, grid [][]int, value, pointX, pointY, maxX, maxY int) {
	if pointX < 0 || pointX >= maxX || pointY < 0 || pointY >= maxY { // check if point exist
		return
	}

	val := grid[pointY][pointX]
	if val == 0 { // check if point actual
		return
	}
	if value+val > *mg { // check if we have more than mg
		(*mg) = value + val
	}
	grid[pointY][pointX] = 0

	//run dfs
	fmt.Println(val, ":", pointX, pointY, ":", value+val)
	dfs(mg, grid, value+val, pointX-1, pointY, maxX, maxY)
	dfs(mg, grid, value+val, pointX+1, pointY, maxX, maxY)
	dfs(mg, grid, value+val, pointX, pointY-1, maxX, maxY)
	dfs(mg, grid, value+val, pointX, pointY+1, maxX, maxY)
}

func getMaximumGold(grid [][]int) int {
	mg := 0
	maxY := len(grid)
	maxX := 0
	if maxY != 0 {
		maxX = len(grid[0])
	}

	for pointY, vY := range grid {
		for pointX, _ := range vY {
			gridCopy := [][]int{}
			for _, v := range grid {
				gridVCopy := []int{}
				for _, vv := range v {
					gridVCopy = append(gridVCopy, vv)
				}
				gridCopy = append(gridCopy, gridVCopy)
			}
			dfs(&mg, gridCopy, 0, pointX, pointY, maxX, maxY)
		}
	}
	return mg
}

func main() {
	//fmt.Println(getMaximumGold([][]int{
	//	[]int{0, 0, 0, 0, 0, 0, 32, 0, 0, 20},
	//	[]int{0, 0, 2, 0, 0, 0, 0, 40, 0, 32},
	//	[]int{13, 20, 36, 0, 0, 0, 20, 0, 0, 0},
	//	[]int{0, 31, 27, 0, 19, 0, 0, 25, 18, 0},
	//	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	//	[]int{0, 0, 0, 0, 0, 0, 0, 18, 0, 6},
	//	[]int{0, 0, 0, 25, 0, 0, 0, 0, 0, 0},
	//	[]int{0, 0, 0, 21, 0, 30, 0, 0, 0, 0},
	//	[]int{19, 10, 0, 0, 34, 0, 2, 0, 0, 27},
	//	[]int{0, 0, 0, 0, 0, 34, 0, 0, 0, 0},
	//}))
}
