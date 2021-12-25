package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	// 按要求只使用O(n)的空间
	rslt := make([]int, rowIndex+1)
	rslt[0] = 1

	if rowIndex == 0 {
		return rslt
	}

	rslt[1] = 1
	if rowIndex == 1 {
		return rslt
	}

	// 逐层计算
	for i:=2; i<=rowIndex; i++ {
		// 每一层从后向前inplace更新
		for j:=i-1; j>0; j-- {
			rslt[j] += rslt[j-1]
		}
		// 末尾放置一个1
		rslt[i] = 1
	}

	return rslt
}

func main() {
	rowIndex := []int{3,0,1}
	for i := range rowIndex {
		fmt.Println(getRow(rowIndex[i]))
	}
}
