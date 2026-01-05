// the row should have only 1 negative val?
// so if there are only 1 negative val i can't do anything... so maybe I should just analyze them?
// i can every 2 of negative turn to be positive

func AbsI(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func maxMatrixSum(matrix [][]int) (res int64) {
	minimal := -1
	even_n := false

	for _, yy := range matrix {
		for _, xx := range yy {
			res += int64(AbsI(xx, 0))
			if AbsI(xx, 0) < minimal || minimal == -1 {
				minimal = AbsI(xx, 0)
			}
			if xx < 0 {
				if even_n {
					even_n = false
				} else {
					even_n = true
				}
			}
		}
	}
	if even_n {
		res -= int64(minimal * 2)
	}

	return
}
