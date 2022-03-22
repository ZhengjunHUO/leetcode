package main

import (
	"fmt"
)

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) > m*n {
		return [][]int{}
	}

	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}

	row, col := 0, 0
	for i := range original {
		ret[row][col] = original[i]
		col++
		if col == n {
			row += 1
			col = 0
		}
	}

	return ret
}

func main() {
	original := [][]int{[]int{1,2,3,4}, []int{1,2,3}, []int{1,2}}
	m,n := []int{2,1,1}, []int{2,3,1}
	for i := range original {
		fmt.Println(construct2DArray(original[i], m[i], n[i]))
	}
}
