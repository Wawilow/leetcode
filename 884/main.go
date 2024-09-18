package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) (res []string) {
	s := strings.Split(s1, " ")
	m := make(map[string]int)
	for _, v := range append(s, strings.Split(s2, " ")...) {
		m[v] = m[v] + 1
	}
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return
}

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}
