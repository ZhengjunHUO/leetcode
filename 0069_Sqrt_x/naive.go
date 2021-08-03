package main

import (
	"fmt"
)

func mySqrt(x int) int {
	l, r := 1, x

	for l <= r {
		m := l + (r-l)/2
		temp := x / m
		switch {
			case m == temp:
				return m
			case m < temp:
				l = m + 1
			case m > temp:
				r = m - 1
		}
	}

	return r
}

func main() {
	xs := []int{4, 8}
	for _,v := range xs {
		fmt.Println(mySqrt(v))
	}
}
