package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 最小面额为1，故最大值不会超过amount/1
	for i := range dp {
		dp[i] = amount+1
	}
	// 确保dp[0]为0, 即amount为0时需要0枚硬币
	dp[0] = 0

	for i:=0; i<=amount; i++ {
		for _,c := range coins {
			if i - c < 0 {
				continue
			}

			// dp[i]即凑到金额i需要的硬币数等于: min{dp[i-c]|c属于coins}+1
			if temp := dp[i-c] + 1; temp < dp[i] {
				dp[i] = temp
			}
		} 
	}

	if dp[len(dp)-1] == amount+1 {
		return -1	
	}else{
		return dp[len(dp)-1]
	}
}

func main() {
	coins := [][]int{[]int{1,2,5}, []int{2}, []int{1}, []int{1}, []int{1}}
	amounts := []int{11,3,0,1,2}

	for i := range coins {
		fmt.Println(coinChange(coins[i], amounts[i]))
	}
}
