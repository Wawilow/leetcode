package main

type coords struct {
	x int
	y int
}

func rectangleDfs(matrix [][]byte, st coords, ed coords) (coords, coords) {
	if ed.x+1 < len(matrix) {
		all_one := true
		for _, v := range matrix[ed.x+1][st.y : ed.y+1] {
			if v != 49 {
				all_one = false
				break
			}
		}
		if all_one {
			ed.x += 1
			return rectangleDfs(matrix, st, ed)
		}
	}
	if st.y-1 >= 0 {
		all_one := true
		for _, v := range matrix[st.x : ed.x+1] {
			if v[st.y-1] != 49 {
				all_one = false
				break
			}
		}
		if all_one {
			st.y -= 1
			return rectangleDfs(matrix, st, ed)
		}
	}
	if len(matrix) > 0 {
		if ed.y+1 < len(matrix[0]) {
			all_one := true
			for _, v := range matrix[st.x : ed.x+1] {
				if v[ed.y+1] != 49 {
					all_one = false
					break
				}
			}
			if all_one {
				ed.y += 1
				return rectangleDfs(matrix, st, ed)
			}
		}
	}
	if st.x-1 >= 0 {
		all_one := true
		for _, v := range matrix[st.x-1][st.y : ed.y+1] {
			if v != 49 {
				all_one = false
				break
			}
		}
		if all_one {
			st.x -= 1
			return rectangleDfs(matrix, st, ed)
		}
	}

	return st, ed
}

func countArea(st, ed coords) int {
	return (ed.x - st.x + 1) * (ed.y - st.y + 1)
}

func maximalRectangle(matrix [][]byte) (res int) {
	for x, xx := range matrix {
		for y, yy := range xx {
			if yy == 49 {
				st, ed := rectangleDfs(matrix, coords{x, y}, coords{x, y})
				fmt.Println(x, y, countArea(st, ed), st, ed)
				if countArea(st, ed) > res {
					res = countArea(st, ed)
				}
			}
		}
	}
	return res
}
