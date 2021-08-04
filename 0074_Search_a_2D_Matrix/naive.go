package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	col := len(matrix[0])
	l, r := 0, len(matrix) * col  - 1

	for l <= r {
		m := l + (r - l) / 2
		mid := matrix[m/col][m%col]
		switch {
		case mid == target:
			return true
		case mid < target:
			l = m + 1
		case mid > target:
			r = m - 1
		}
	}

	return false
}

func main() {
	matrix := [][]int{[]int{1,3,5,7},[]int{10,11,16,20},[]int{23,30,34,60}}
	targets := []int{3, 13}

	for _, v := range targets {
		fmt.Println(searchMatrix(matrix, v))
	}
}
