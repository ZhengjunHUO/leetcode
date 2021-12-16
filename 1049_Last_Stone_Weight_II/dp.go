package main

import (
	"fmt"
)

func max (a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
  类似0494_Target_Sum，可以在每个数前加上正号或负号，分成两组
    A + B = sum(stones)
    A + (-B) = target
    且A>=B
  得到target = sum - 2B, 需要最小化target则需要最大化B, B <= SUM/2
  确定背包容量上限为SUM/2，求最大能装多少重量的石头
*/
func lastStoneWeightII(stones []int) int {
	sum := 0	
	for i := range stones {
		sum += stones[i]
	}
	limit := sum/2 + 1

	dp := make([]int, limit)
	
	for _,v := range stones {
		for j:=len(dp)-1; j>=0; j-- {
			if v <= j {
				dp[j] = max(dp[j], dp[j-v]+v)
			}
		}
	}

	return sum - dp[len(dp)-1] * 2
}

func main() {
	stones := [][]int{[]int{2,7,4,1,8,1}, []int{31,26,33,21,40}, []int{1,2}}
	for i := range stones {
		fmt.Println(lastStoneWeightII(stones[i]))
	}
}
