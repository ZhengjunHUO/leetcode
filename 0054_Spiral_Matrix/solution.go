package main

import (
	"fmt"
)

// 学习了@stellari的方案
func spiralOrder(matrix [][]int) []int {
	// 向右，向下，向左，向上 每一步对matrix坐标的增减值 eg. (0, -1)开始向右一格为+(0, 1) => (0, 0) 
	directions := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	// 记录当前的前进方向，在遍历完一行或一列后递增(模len(directions))
	idxDir := 0

	// 矩阵行数，列数
	rows, cols := len(matrix), len(matrix[0]) 
	rslt := make([]int, 0, rows * cols)
	// 从(0, -1)出发，每次遍历行(或列)的长度 eg. 3*3矩阵为右3下2左2上1右1下0
	steps := []int{cols, rows-1}	
	idxR, idxC := 0, -1	

	// 每完成行或列的遍历steps中相应的值会递减，已确定下一次行或列遍历的长度，直至0 (见上3*3的例子)
	for steps[idxDir%2] != 0 {
		for i := 0; i < steps[idxDir%2]; i++ {
			idxR, idxC = idxR + directions[idxDir][0], idxC + directions[idxDir][1]
			rslt = append(rslt, matrix[idxR][idxC])
		}		
		steps[idxDir%2] -= 1
		idxDir = (idxDir + 1)%4	
	} 

	return rslt
}

func main() {
	matrix := [][][]int{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}, [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}}
	for i := range matrix {
		fmt.Println(spiralOrder(matrix[i]))
	}
}
