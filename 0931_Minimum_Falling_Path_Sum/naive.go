package main

import (
	"fmt"
)

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	temp := min2(a, b)
	return min2(temp, c)
}

func minFallingPathSum(matrix [][]int) int {
	dp, prev := make([]int, len(matrix[0])), make([]int, len(matrix[0]))
	// 可以从第一行任意一点出发
	copy(dp, matrix[0])

	for i:=1; i<len(matrix); i++ {
		// 下一行最小值的计算基于dp但也同时会覆盖dp，需要一个dp的拷贝避免影响
		copy(prev, dp)
		for j := range matrix[i] {
			// 根据题目给出的规则可知 dp[x] = min(dp[x-1], dp[x], dp[x+1]) + matrix[j][x]
			if j == 0 {
				dp[j] = min2(prev[j],prev[j+1])+matrix[i][j]
				continue
			}
			if j == len(matrix[i]) - 1 {
				dp[j] = min2(prev[j-1], prev[j])+matrix[i][j]
				continue
			}

			dp[j] = min3(prev[j-1], prev[j], prev[j+1])+matrix[i][j]
		}
	}

	rslt := 10001
	for i := range dp {
		if dp[i] < rslt {
			rslt = dp[i]
		}
	}
	return rslt
}

func main() {
	matrixs := [][][]int{[][]int{[]int{2,1,3},[]int{6,5,4},[]int{7,8,9}}, [][]int{[]int{-19,57},[]int{-40,-5}}, [][]int{[]int{-48}}}
	for i := range matrixs {
		fmt.Println(minFallingPathSum(matrixs[i]))
	}
}
