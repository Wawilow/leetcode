package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	memory := make(map[string]string)
	sPattert := strings.Split(pattern, "")
	sS := strings.Split(s, " ")
	if len(sPattert) != len(sS) {
		return false
	}
	for i, _ := range sPattert {
		_, exits := memory[sPattert[i]]
		if !exits {
			if !alredyIn(memory, sS[i]) {
				memory[sPattert[i]] = sS[i]
			} else {
				return false
			}
		}
		fmt.Println(memory[sPattert[i]], sPattert[i], sS[i])
		if memory[sPattert[i]] != sS[i] {
			return false
		}
	}
	return true
}

func alredyIn(memory map[string]string, v string) bool {
	for _, val := range memory {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
