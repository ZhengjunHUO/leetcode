package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([]int, n)
	/*
	  dp[i]表示进入该格子时至少需要的生命值
	  dp在压缩前按二维矩阵表示，从下往上，从右往左计算，最后返回dp[0][0]
	  关键是如果计算得到负值表示迷宫格子中是个血瓶，进入所需最小生命值以1计算
		x x 2	  7 5  2
		x x 5  => 6 11 5
		1 1 6	  1 1  6
	*/
	dp[n-1] = max(1, 1 - dungeon[m-1][n-1])
	for i := n-2; i >= 0; i-- {
		dp[i] = max(1, dp[i+1] - dungeon[m-1][i])
	}

	for j := m-2; j >= 0; j-- {
		dp[n-1] = max(1, dp[n-1] - dungeon[j][n-1])
		for i := n-2; i >= 0; i-- {
			dp[i] = min(max(1,dp[i]-dungeon[j][i]), max(1,dp[i+1]-dungeon[j][i]))
		}
	}

	return dp[0]
}

func main() {
	ds := [][][]int{[][]int{[]int{-2,-3,3},[]int{-5,-10,1},[]int{10,30,-5}}, [][]int{[]int{0}}}
	for i := range ds {
		fmt.Println(calculateMinimumHP(ds[i]))
	}
}
