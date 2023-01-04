package main

import "fmt"

func alphasort(str string) bool {
	strNums := []int32{}
	for _, v := range str {
		strNums = append(strNums, v)
	}
	for i := 0; i < len(str)-1; i++ {
		for j := 0; j < len(str)-i-1; j++ {
			if strNums[j] > strNums[j+1] {
				return false
				//newStr[j], newStr[j+1] = newStr[j+1], newStr[j]
				//fmt.Println(newStr, j, j+1)
			}
		}
	}
	return true
}

func invertAlphasort(str string) bool {
	strNums := []int32{}
	for _, v := range str {
		strNums = append(strNums, v)
	}
	for i := 0; i < len(str)-1; i++ {
		for j := 0; j < len(str)-i-1; j++ {
			if strNums[j] < strNums[j+1] {
				return false
				//newStr[j], newStr[j+1] = newStr[j+1], newStr[j]
				//fmt.Println(newStr, j, j+1)
			}
		}
	}
	return true
}

func minDeletionSize(strs []string) int {
	res := 0
	for _, v := range strs {
		if !alphasort(v) {
			if !invertAlphasort(v) {
				//fmt.Println(v)
				res += 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
	//fmt.Println(invertAlphasort("cba"))
}
