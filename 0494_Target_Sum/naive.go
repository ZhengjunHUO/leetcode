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
				/* 对2维DP数组来说写成：
				  dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
				  1) 不把ith物品装入背包那么其方法数就继承dp[i-1][j]
				  2) 把ith物品装入了背包，就要看前i-1个物品有几种方法可以装满j-nums[i-1]的容量
				  最后结果取(1)+(2)的和
				*/
				dp[j] = dp[j] + dp[j-nums[i]]
			}
			// else的话dp[j] = dp[j] 可省略(dp[i][j] = dp[i-1][j])
		}
	}

	return dp[len(dp)-1]
}

func main() {
	/*
	[1 1 0 0 0]
	[1 2 1 0 0]
	[1 3 3 1 0]
	[1 4 6 4 1]
	[1 5 10 10 5]
	*/
	fmt.Println(findTargetSumWays([]int{1,1,1,1,1}, 3))
}
