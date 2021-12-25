package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	rslt := make([][]int, 0, numRows)
	rslt = append(rslt, []int{1})

	// numRows为1直接返回第一行
	if numRows == 1 {
		return rslt
	}

	// numRows为2直接返回前两行
	rslt = append(rslt, []int{1, 1})
	if numRows == 2 {
		return rslt
	}

	for i:=2; i<numRows; i++ {
		// 准备当前行，先插入一个1
		line := make([]int, 0, i+1)
		line = append(line, 1)

		// 插入上一行的index[0, 1]的和 直到 index[i-2, i-1]的和 
		for j:=1; j<i; j++ {
			line = append(line, rslt[i-1][j-1]+rslt[i-1][j])
		}
		// 末尾插入1
		line = append(line, 1)
		
		// 将当前行插入结果
		rslt = append(rslt, line)
	}

	return rslt
}

func main() {
	numRows := []int{5,1,9}
	for i := range numRows {
		fmt.Println(generate(numRows[i]))
	}
}
