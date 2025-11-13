func maxOperations(s string) (res int) {
    counter := 0
    leftOnes := 0
    for _, j := range s {
		if j == 49 {counter += 1} else if counter != 0{
			res += leftOnes + counter
            leftOnes += counter
            counter = 0
		}}
    return res}