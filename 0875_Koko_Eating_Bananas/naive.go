package main

import (
	"fmt"
)

func timeNeeded(piles []int, x int) int {
	rslt := 0
	for i := range piles {
		rslt += piles[i]/x

		if piles[i]%x != 0 {
			rslt++
		}
	}

	return rslt
}  

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 1000000000+1
	for l < r {
		m := l + (r-l)/2
		if timeNeeded(piles, m) <= h {
			r = m
		}else{
			l = m + 1
		}
	}

	return l
}

func main() {
	ps := [][]int{[]int{3,6,7,11}, []int{30,11,23,4,20}, []int{30,11,23,4,20}}
	hs := []int{8, 5, 6}

	for i := range ps {
		fmt.Println(minEatingSpeed(ps[i], hs[i]))
	}
}
