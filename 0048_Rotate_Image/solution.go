package main

import (
	"fmt"
)

// 顺时针转动90度
func rotate(matrix [][]int)  {
	// 上下翻转 （如果逆时针转动的话需要左右翻转）
	for i, j := 0, len(matrix) - 1; i < j ; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}    

	// 轴对称对调元素
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][][]int{[][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}, [][]int{[]int{5,1,9,11},[]int{2,4,8,10},[]int{13,3,6,7}, []int{15,14,12,16}}, [][]int{[]int{1}}, [][]int{[]int{1,2}, []int{3,4}}}
	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		rotate(matrix[i])
	}
	fmt.Println(matrix)
}
