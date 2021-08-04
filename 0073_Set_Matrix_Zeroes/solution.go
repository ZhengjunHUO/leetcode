package main

import (
	"fmt"
)

/* 
根据Hint给出的提示，参考@mzchen的简练代码
每行第一个元素记录该行是否需要置0
每列第一个元素记录该列是否需要置0
第一行第一列需要额外的一个格子来存储该列的信息
*/
func setZeroes(matrix [][]int)  {
	firstCol, row, col := 1, len(matrix), len(matrix[0])

	for i:=0; i<row; i++ {
		// 单独处理第一列
		if matrix[i][0] == 0 {
			firstCol = 0
		}
	
		// 从第二列开始
		for j:=1; j<col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i:=row-1; i>=0; i-- {
		for j:=col-1; j>=1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if firstCol == 0 {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][][]int{[][]int{[]int{1,1,1},[]int{1,0,1},[]int{1,1,1}}, [][]int{[]int{0,1,2,0},[]int{3,4,5,2},[]int{1,3,1,5}}}
	for i := range matrix {
		setZeroes(matrix[i])
		fmt.Println(matrix[i])
	}
}
