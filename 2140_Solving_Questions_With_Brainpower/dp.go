package main

import (
	"fmt"
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func mostPoints(questions [][]int) int64 {
	// 对于每个元素i，记录从i开始可以得到多少点
	dp := make([]int64, len(questions))
	// nextAvailable：当前跳过brainpower个元素后的位置开始可以获得的点数
	// next：下一个位置开始可以获得的点数
	var nextAvailable, next int64
	for i := len(dp) - 1; i >= 0 ; i-- {
		// 出界的index点数为0
		if nextAvailableIdx := i + questions[i][1] + 1; nextAvailableIdx >= len(dp) {
			nextAvailable = 0
		}else{
			nextAvailable = dp[nextAvailableIdx]
		}

		// 出界的index点数为0
		if i + 1 >= len(dp) {
			next = 0
		}else{
			next = dp[i+1]
		}

		// 在(不做这题，做题)获得的点数之间取大值 
		dp[i] = max(next, int64(questions[i][0])+nextAvailable)
	}

	return dp[0]
}

func main() {
	qs := [][][]int{[][]int{[]int{3,2},[]int{4,3},[]int{4,4},[]int{2,5}}, [][]int{[]int{1,1},[]int{2,2},[]int{3,3},[]int{4,4},[]int{5,5}}}
	for i := range qs {
		fmt.Println(mostPoints(qs[i]))
	}
}
