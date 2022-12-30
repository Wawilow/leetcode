package main

func floor(n float64) int {
	if n-float64(int(n)) > 0 {
		return int(n) + 1
	} else {
		return int(n)
	}
}

func minStoneSum(piles []int, k int) int {
	if k > len(piles) {
		k = len(piles)
	}

	afterPiles := make([]int, len(piles))
	for i, v := range piles {
		afterPiles[i] = floor(float64(v) / 2.0)
	}
	for i := 0; i < len(afterPiles)-1; i++ {
		for j := 0; j < len(afterPiles)-i-1; j++ {
			if afterPiles[j] < afterPiles[j+1] {
				afterPiles[j], afterPiles[j+1] = afterPiles[j+1], afterPiles[j]
				piles[j], piles[j+1] = piles[j+1], piles[j]
			}
		}
	}
	for i := 0; i < k; i++ {
		piles[i] = afterPiles[i]
	}

	res := 0
	for _, v := range piles {
		res += v
	}
	return res
}

//func main() {
//	minStoneSum([]int{5, 4, 9}, 2)
//fmt.Println(math.Floor(2.6))
//}
