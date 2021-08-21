package main

import (
	"fmt"
)

func daysNeeded(weights []int, capacity int) int {
	rest, rslt := capacity, 1
	for i := range weights {
		if weights[i] > capacity {
			return 50001
		}

		if weights[i] > rest {
			rslt++
			rest = capacity - weights[i]
		}else{
			rest = rest - weights[i]
		}
	}

	return rslt
}

func shipWithinDays(weights []int, days int) int {
	l, r := 1, 25000000+1	
	for l < r {
		m := l + (r-l)/2
		if daysNeeded(weights, m) <= days {
			r = m
		}else{
			l = m + 1
		}
	}

	return l
}

func main() {
	ws := [][]int{[]int{1,2,3,4,5,6,7,8,9,10}, []int{3,2,2,4,1,4}, []int{1,2,3,1,1}}
	ds := []int{5, 3, 4}

	for i := range ws {
		fmt.Println(shipWithinDays(ws[i], ds[i]))
	}
}
