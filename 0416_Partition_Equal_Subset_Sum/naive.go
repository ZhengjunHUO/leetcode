package main

import (
	"fmt"
)

// 思路和0494_Target_Sum共通，将分成两组的问题转化为背包问题
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	
	// 如果和是奇数不可能分成和相同的两组
	if sum%2 == 1 {
		return false
	}

	// 遍历所有元素看能否挑选出组合填满背包（总和的一半）
	// 状态为背包容量大小，base case为容量等于0时默认可以填满
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true

	for i := range nums {
		// dp表压缩为一维数组，所以需要从后往前更新
		for j := sum; j >= 1; j-- {
			if nums[i] <= j {
				// 加入当前元素看是否可以填满 或 不加入当前元素（继承上个元素结果）
				// j-nums[i]如果正好等于0表示当前元素正好填满了背包可以从dp[0]处获得一个true
				dp[j] = dp[j-nums[i]] || dp[j]
			}
			// 如果nums[i] > j即不可取，只能继承前一个元素的结果，dp[j] = dp[j]可省略
		}
	}	

	return dp[len(dp)-1]
}

func main() {
	nums := [][]int{[]int{1,5,11,5}, []int{1,2,3,5}}
	for i := range nums {
		fmt.Println(canPartition(nums[i]))
	}
}
