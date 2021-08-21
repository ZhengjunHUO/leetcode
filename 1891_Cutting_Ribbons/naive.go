package main

import (
	"fmt"
)

func getNum(ribbons []int, max int) int {
	rslt := 0
	for i := range ribbons {
		rslt += ribbons[i]/max
	}
	return rslt
}

func maxLength(ribbons []int, k int) int {
	l, r := 1, 0
	for i := range ribbons {
		if ribbons[i] > r {
			r = ribbons[i]
		}
	}


	for l < r {
		m := l + (r - l)/2
		if getNum(ribbons, m) >= k {
			l = m + 1
		}else{
			r = m
		}
	}

	return l-1
}

func main() {
	rbs := [][]int{[]int{9,7,5}, []int{7,5,9}, []int{5,7,9}}
	ks := []int{3,4,22}

	for i := range rbs {
		fmt.Println(maxLength(rbs[i], ks[i]))
	}
}
