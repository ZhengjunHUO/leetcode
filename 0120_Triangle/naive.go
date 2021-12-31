package main

import (
	"fmt"
)

const MaxInt = int(^uint(0)>> 1)

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	// 用来存放对每一行计算的中间结果
	table := make([]int, n)   
	for i := range table {
		table[i] = MaxInt
	}

	// 没有必要计算最后一行
	for i:=0; i<n-1; i++ {
		for j:=0; j<len(triangle[i]); j++ {
			// 计算到下一层两个邻接节点的花费，取最小值存放于缓存中
			table[j] = Min(table[j], triangle[i][j]+triangle[i+1][j])
			table[j+1] = Min(table[j+1], triangle[i][j]+triangle[i+1][j+1])
		}

		// 计算完当前层后把缓存的数据写入triangle中，同时初始化缓存
		for k:=0; k<=i+1; k++ {
			triangle[i+1][k] = table[k]
			table[k] = MaxInt
		}
	}

	fmt.Println(triangle)

	min := MaxInt
	for i:=0; i<len(triangle[n-1]); i++ {
		min = Min(min, triangle[n-1][i])
	}

	return min
}
