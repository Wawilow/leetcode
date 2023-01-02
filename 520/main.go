package main

import "strings"

func detectCapitalUse(word string) bool {
	if strings.ToLower(word) == word {
		return true
	}
	if strings.ToTitle(word) == word {
		return true
	}
	count := 0
	for _, v := range strings.Split(word, "") {
		//v := word[i]
		if strings.ToTitle(v) == v {
			count += 1
		}
	}
	if count == 1 && strings.ToTitle(strings.Split(word, "")[0]) == strings.Split(word, "")[0] {
		return true
	}
	return false
}

func main() {

}
