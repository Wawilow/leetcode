package main

import "fmt"

func main() {
	var s = "MCMXCIV"

	const I = "I"
	const V = "V"
	const X = "X"
	const L = "L"
	const C = "C"
	const D = "D"
	const M = "M"
	whileRun := true
	parse := map[string]string{
		"I": "",
		"V": "",
		"X": "",
		"L": "",
		"C": "",
		"D": "",
		"M": "",
	}
	res := 0
	for whileRun {
		switch {
		case parse[I] == "":
			if string(s[len(s)-1]) != I {
				parse[I] = "-1"
			}
			for i := len(s) - 1; i >= 0; i-- {
				if string(s[i]) != I {
					break
				}
				parse[I] = string(s[i]) + parse[I]
				res += 1
				s = s[:i]
			}
		case parse[V] == "":
			if string(s[len(s)-1]) != V {
				parse[V] = "-1"
			}
			for i := len(s) - 1; i >= 0; i-- {
				if string(s[i]) != V && string(s[i]) != I {
					break
				}
				parse[V] = string(s[i]) + parse[V]
				switch {
				case string(s[i]) == V:
					res += 5
				case string(s[i]) == I:
					res -= 1
				}
				s = s[:i]
			}
		case parse[X] == "":
			if string(s[len(s)-1]) != X {
				parse[X] = "-1"
			}
			for i := len(s) - 1; i >= 0; i-- {
				if string(s[i]) != X && string(s[i]) != I {
					break
				}
				parse[X] = string(s[i]) + parse[X]
				switch {
				case string(s[i]) == X:
					res += 10
				case string(s[i]) == I:
					res -= 1
				}
				s = s[:i]
			}
		case parse[L] == "":
			if string(s[len(s)-1]) != L {
				parse[L] = "-1"
			}
			for i := len(s) - 1; i >= 0; i-- {
				if string(s[i]) != L && string(s[i]) != X {
					break
				}
				parse[L] = string(s[i]) + parse[L]
				switch {
				case string(s[i]) == L:
					res += 50
				case string(s[i]) == X:
					res -= 10
				}
				s = s[:i]
			}
		case parse[C] == "":
			if string(s[len(s)-1]) != C {
				parse[C] = "-1"
			}
			for i := len(s) - 1; i >= 0; i-- {
				if string(s[i]) != C && string(s[i]) != X {
					break
				}
				parse[C] = string(s[i]) + parse[C]
				switch {
				case string(s[i]) == C:
					res += 100
				case string(s[i]) == X:
					res -= 10
				}
				s = s[:i]
			}
		case parse[D] == "":
			if string(s[len(s)-1]) != D {
				parse[D] = "-1"
			}
			for i := len(s) - 1; i >= 0; i-- {
				if string(s[i]) != D && string(s[i]) != C {
					break
				}
				parse[D] = string(s[i]) + parse[D]
				switch {
				case string(s[i]) == D:
					res += 500
				case string(s[i]) == C:
					res -= 100
				}
				s = s[:i]
			}
		case parse[M] == "":
			if string(s[len(s)-1]) != M {
				parse[M] = "-1"
			}
			for i := len(s) - 1; i >= 0; i-- {
				if string(s[i]) != M && string(s[i]) != C {
					break
				}
				parse[M] = string(s[i]) + parse[M]
				switch {
				case string(s[i]) == M:
					res += 1000
				case string(s[i]) == C:
					res -= 100
				}
				s = s[:i]
			}
		default:
			whileRun = false
		}
		if s == "" {
			whileRun = false
		}
	}

	fmt.Println("Res ", res)
}
