package main

import "fmt"

func uniquePathsIII(grid [][]int) int {
	res := 0
	path := [][]int{}
	pos := ppos(grid)
	n := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				n += 1
			}
		}
	}
	dfs(&res, grid, path, pos, n)

	return res
}

func dfs(res *int, grid [][]int, path [][]int, pos []int, n int) {
	if valid(grid, path, pos) {
		if grid[pos[0]][pos[1]] != 2 {
			path = append(path, pos)
			//fmt.Println("---", path, pos)
			dfs(res, grid, path, []int{pos[0] - 1, pos[1]}, n)
			dfs(res, grid, path, []int{pos[0] + 1, pos[1]}, n)
			dfs(res, grid, path, []int{pos[0], pos[1] - 1}, n)
			dfs(res, grid, path, []int{pos[0], pos[1] + 1}, n)
		} else {
			if len(path)+1 >= n+2 {
				fmt.Println(path, pos)
				*res += 1
			}
		}
	}
}

func ppos(grid [][]int) []int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func valid(grid [][]int, path [][]int, pos []int) bool {
	hight, lenght := len(grid), len(grid[0])
	if 0 <= pos[0] && pos[0] < hight && 0 <= pos[1] && pos[1] < lenght {
		if grid[pos[0]][pos[1]] == -1 {
			return false
		} else {
			res := true
			for _, v := range path {
				//fmt.Println("v: ", ((v[0] == pos[0]) && (v[1] == pos[1])))
				if (v[0] == pos[0]) && (v[1] == pos[1]) {
					res = false
				}
			}
			return res
		}
	}
	return false
}

func main() {
	//grid := [][]int{[]int{1, 0, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 2, -1}}
	grid := [][]int{[]int{0, 1}, []int{2, 0}}
	//fmt.Println(valid(grid, [][]int{[]int{1, 1}}, []int{1, 1}))
	fmt.Println(uniquePathsIII(grid))
}
