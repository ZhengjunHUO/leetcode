package main

import (
	"fmt"
)

/*
把得到正号和负号的两组分开，可得sum(pos) - sum(neg) = target
两边同时加上sum(all)，转化为2sum(pos) = target + sum(all)
问题变成挑选元素使其和等于(target + sum(all))/2，即背包问题
*/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum < target || (sum + target)%2 == 1 {
		return 0
	}

	goal := (sum + target)/2

	dp := make([]int, goal+1)
	// 背包容量为0，只有不放1种选择
	dp[0] = 1

	// 对于每个元素，从后向前计算dp表
	for i := range nums {
		for j:=goal; j>=1; j-- {
			if j >= nums[i] {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
			// else的话dp[j] = dp[j]
		}
	}

	return dp[len(dp)-1]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1,1,1,1,1}, 3))
}
