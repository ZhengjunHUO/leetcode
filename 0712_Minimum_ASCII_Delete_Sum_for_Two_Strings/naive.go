package main

import (
	"fmt"
)

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else{
				/* 
				  左: 删除s1的当前字母以到达前一个平衡dp[i-1][j]，花费为s1当前字母的值加上继承的前一个平衡的花费   
				  右: 删除s2的当前字母，逻辑同上

				  如sea和eat的dp矩阵
				  	0   101 198 314
					115 216 313 429
					216 115 212 328
					313 212 115 231
				  最后一格dp[3][3]的计算: min(328+97, 115+116)  [97: a, 116: t]
				  左侧选择删除a(97)，此时为"se"对"eat"，对应的dp[i-1][j]为328意味着需要到达等价需要付出删掉s,a,t的代价
				  右侧选择删除t(116)，此时为"sea"对"ea"，对应的dp[i][j-1]为115意味着需要到达等价只需要删掉s(115)即可
				*/
				dp[i][j] = min2((dp[i-1][j] + int(s1[i-1])), (dp[i][j-1] + int(s2[j-1])))
			}
		}
	}

	return dp[m][n]
}

func main() {
	s1 := []string{"sea", "delete"}
	s2 := []string{"eat", "leet"}

	for i := range s1 {
		fmt.Println(minimumDeleteSum(s1[i], s2[i]))
	}
}
