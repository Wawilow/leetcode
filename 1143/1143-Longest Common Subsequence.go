package _143

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func delCharter(text string, id int) string {
	res := ""
	if id == 0 {
		res = text[1:]
	} else if id == len(text)-1 {
		res = text[:len(text)-1]
	} else {
		res = text[:id] + text[id+1:]
	}
	return res
}

func subsequences(res *[]string, text string, howMuch int) *[]string {
	if howMuch == 0 {
		*res = append(*res, text)
		return res
	}

	for i := 0; i < len(text); i++ {
		if howMuch == 1 {
			*res = append(*res, delCharter(text, i))
		} else {
			subsequences(res, delCharter(text, i), howMuch-1)
		}
		//fmt.Println(i)
		//if i == 0 {}
	}
	return res
}

func longestCommonSubsequence(text1 string, text2 string) int {
	length := 0
	if len(text1) >= len(text2) {
		length = len(text2)
		for i := 0; i < length; i++ {
			text1Subsequences := []string{}
			subsequences(&text1Subsequences, text1, i+(len(text1)-len(text2)))
			text2Subsequences := []string{}
			subsequences(&text2Subsequences, text2, i)
			for _, sequence := range text1Subsequences {
				if contains(text2Subsequences, sequence) {
					return len(sequence)
				}
			}
		}
	} else {
		length = len(text1)
		for i := 0; i < length; i++ {
			text1Subsequences := []string{}
			subsequences(&text1Subsequences, text1, i)
			text2Subsequences := []string{}
			subsequences(&text2Subsequences, text2, i+(len(text2)-len(text1)))
			for _, sequence := range text1Subsequences {
				if contains(text2Subsequences, sequence) {
					return len(sequence)
				}
			}
		}
	}
	//for i := 0; i < length; i++ {
	//	text1Subsequences := []string{}
	//	subsequences(&text1Subsequences, text1, len(text1))
	//	text2Subsequences := []string{}
	//	subsequences(&text2Subsequences, text2, i)
	//	fmt.Println(i, text1Subsequences, text2Subsequences)
	//	for _, sequence := range text1Subsequences {
	//		if contains(text2Subsequences, sequence) {
	//			return len(sequence)
	//		}
	//	}
	//}
	return 0
}

//func main() {
//	fmt.Println(longestCommonSubsequence("bsbininm", "jmjkbkjkv"))
//	//text1Subsequences := []string{}
//	//subsequences(&text1Subsequences, "abcdef", 3)
//	//fmt.Println(text1Subsequences)
//}
