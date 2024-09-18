package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	h int
	m int
}

func timeDifference(t1, t2 Time) (r int) {
	if t1.h == t2.h {
		if t1.m > t2.m {
			return (24 * 60) - t1.m + t2.m
		} else {
			return t2.m - t1.m
		}
	}

	if t1.h > t2.h {
		r += (24 - t1.h) + t2.h
	} else {
		r += t2.h - t1.h
	}
	r *= 60

	return r + t2.m - t1.m
}

func findMinDifference(timePoints []string) (r int) {
	l := []Time{}
	for _, t := range timePoints {
		raw := strings.Split(t, ":")
		h, _ := strconv.Atoi(raw[0])
		m, _ := strconv.Atoi(raw[1])
		l = append(l, Time{h: h, m: m})
	}

	r = -1
	for iNum, i := range l {
		for jNum, j := range l {
			if iNum == jNum {
				continue
			}
			dij := timeDifference(i, j)
			fmt.Println(dij, i, j)
			if r == -1 {
				r = dij
			} else if r > dij {
				r = dij
			}
		}
	}
	return
}

func main() {
	//fmt.Println(timeDifference(Time{12, 11}, Time{13, 13}))
	//fmt.Println(timeDifference(Time{23, 1}, Time{1, 11}))
	//fmt.Println(timeDifference(Time{1, 50}, Time{23, 11}))
	fmt.Println(timeDifference(Time{3, 1}, Time{3, 0}))
	fmt.Println(timeDifference(Time{3, 1}, Time{3, 10}))
	//fmt.Println(findMinDifference([]string{"01:01", "02:01", "03:00", "03:01"}))
}
