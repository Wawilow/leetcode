package main

import "fmt"

func convert(s string, numRows int) string {
	zero := "±"

	matr := [][]string{}
	sss := []string{}
	for i, _ := range s {
		sss = append(sss, string(s[i]))
	}
	for i := 0; i < numRows; i++ {
		ss := []string{}
		for j := 0; j < len(s)/numRows+((numRows-2)*(len(s)/numRows)); j++ {
			ss = append(ss, zero)
		}
		matr = append(matr, ss)
	}
	for i := 0; i <= (len(s)/numRows + ((numRows - 2) * (len(s) / numRows))); i++ {
		count := 0
		for l := 0; l < numRows; l++ {
			if count < len(sss) {
				matr[l][i] = sss[count]
				count += 1
			}
		}
		for l := numRows - 1; l > 1; l-- {
			if count < len(sss) {
				i++
				matr[l-1][i] = sss[count]
				count += 1
			}
		}
		sss = sss[count:]
		count = 0
	}
	//fmt.Println(matr)
	for _, v := range matr {
		fmt.Println(v)
	}
	ans := ""
	for _, v := range matr {
		for _, vv := range v {
			if vv != zero {
				ans = fmt.Sprintf("%s%s", ans, vv)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(convert("abcdefghi111111", 4))
	//fmt.Println("PAHNAPLSIIGYIR" == "PAHNAPLSIIGYIR")
}
