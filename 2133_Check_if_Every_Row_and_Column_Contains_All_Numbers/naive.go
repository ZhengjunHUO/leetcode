package main

import (
	"fmt"
)

func checkValid(matrix [][]int) bool {
	// 模拟set
	var col, row map[int]struct{}
	for i := range matrix {
		col, row = map[int]struct{}{}, map[int]struct{}{}
		// 同时检查第i行和第j列，将matrix中的元素加入set
		// 由于1 <= matrix[i][j] <= n，没必要检查元素的具体值
		// 如果符合条件set的大小必定等于n
		for j := range matrix[0] {
			row[matrix[i][j]] = struct{}{}
			col[matrix[j][i]] = struct{}{}
		}

		// 如果set大小不等于n(小于)则不成立
		if len(row) != len(matrix) || len(col) != len(matrix) {
			return false
		}
	}

	return true
}

func main() {
	matrix := [][][]int{[][]int{[]int{1,2,3},[]int{3,1,2},[]int{2,3,1}}, [][]int{[]int{1,1,1},[]int{1,2,3},[]int{1,2,3}}}
	for i := range matrix {
		fmt.Println(checkValid(matrix[i]))
	}
}
