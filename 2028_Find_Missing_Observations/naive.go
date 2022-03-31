package main

import (
	"fmt"
)

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	total := mean * (m+n)

	for i := range rolls {
		total -= rolls[i]
	}

	// can't satisfied
	if total/n > 6 || ( total/n == 6 && total%n != 0) {
		return []int{}
	}

	// distribute points
	base := total/n
	ret := make([]int, n)
	for i := range ret {
		ret[i] = base
	}

	total %= n
	margin := 6 - base
	var idx int
	for total > 0 {
		if total <= margin {
			ret[idx] += total
			break
		}

		ret[idx] = 6
		idx++
		total -= margin
	}

	return ret
}

func main() {
	rolls := [][]int{[]int{3,2,4,3}, []int{1,5,6}, []int{1,2,3,4}}
	means := []int{4,3,6}
	n := []int{2,4,4}

	for i := range n {
		fmt.Println(missingRolls(rolls[i], means[i], n[i]))
	}
}
