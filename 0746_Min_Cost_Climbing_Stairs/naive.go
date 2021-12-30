package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	// rslt[i]：为了爬到i累计消耗的精力，初值为rslt[0] = 0, rslt[1] = 0
	rslt := make([]int, n+1)

	for i:=2; i<n+1; i++ {
		// 取 从i-1出发加上i-1的cost 和 从i-2出发加上i-2的cost 的最小值
		rslt[i] = min(rslt[i-1]+cost[i-1], rslt[i-2]+cost[i-2])
	}

	fmt.Println("[Debug]rslt: ", rslt)
	return rslt[n]
}
