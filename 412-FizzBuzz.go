package main

import "fmt"

func fizzBuzz(n int) []string {
	var ans []string
	for i := 1; i < n+1; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			ans = append(ans, "FizzBuzz")
		case i%3 == 0:
			ans = append(ans, "Fizz")
		case i%5 == 0:
			ans = append(ans, "Buzz")
		default:
			ans = append(ans, fmt.Sprintf("%d", i))
		}
	}
	return ans
}

//func main() {
//	fmt.Print(fizzBuzz(3))
//}
