func largestLocal(grid [][]int) [][]int {
	r := [][]int{}
	rv := []int{}
	n := len(grid)
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			l := 0
			for _, nL := range grid[i : i+3] {
				for _, nLL := range nL[j : j+3] {
					if nLL > l {
						l = nLL
					}
				}
			}
			rv = append(rv, l)
			if len(rv) == n-2 {
				r = append(r, rv)
				rv = []int{}
			}
		}
	}
	return r
}
