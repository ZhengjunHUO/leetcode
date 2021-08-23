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

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	// 初始默认长度为自己，即1
	for i := range dp {
		dp[i] = 1
	}

	for i := range nums {
		for j := 0; j < i; j++ {
			// 遍历前面的元素，如果当前元素大于某个元素，在他的长度上加1，最后保留最大的
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}	
		}
	}

	// 返回dp数列中的最大值
	rslt := 0
	for i := range dp {
		if dp[i] > rslt {
			rslt = dp[i]
		}
	}
	
	return rslt
}

func main() {
	nums := [][]int{[]int{10,9,2,5,3,7,101,18}, []int{0,1,0,3,2,3}, []int{7,7,7,7,7,7,7}}	
	for i := range nums {
		fmt.Println(lengthOfLIS(nums[i]))
	}
}
