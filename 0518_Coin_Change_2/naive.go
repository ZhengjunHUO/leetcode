package main

import (
	"fmt"
)

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	// amount为0，只有一种组合即什么都不取
	dp[0] = 1

	for i := range coins {
		for j := 1; j < len(dp); j++ {
			/*
			  如果使用二维dp表示，相关的转移逻辑为
			  dp[i][j] = d[i-1][j]   (if c > j)
			           = d[i-1][j] + dp[i][j-c]  (if c <= j)
			  j表示当前amount的值，i表示考虑“前i种”硬币有几种组合方法
			  故计算dp[i][j]需要加上历史数值dp[i-1][j]
			*/
			if coins[i] <= j {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}

	return dp[len(dp)-1]
}

func main() {
	/* 1,2,5三种面额硬币凑5的dp
	[1 0 0 0 0 0]	=> 初值
	[1 1 1 1 1 1]	=> 只有面值为1的硬币对应各种amount的方法只有1种
	[1 1 2 2 3 3]   => 1  0+1  1+1  1+1  2+1  2+1 面值为2的硬币从amount=2需要将dp[i][j-2]和上面的格子相加
	[1 1 2 2 3 4]   => 1  0+1  0+2  0+2  0+3  1+3 
	*/
	coins := [][]int{[]int{1,2,5}, []int{2}, []int{10}}
	as := []int{5,3,10}

	for i := range coins {
		fmt.Println(change(as[i], coins[i]))
	}
}
