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

// 计算s的ith字符到jth字符之间的最长回文值
// 如果s[i]正好等于s[j]那么两者都取，其值为s[i+1][j-1]+2
// 否则只取其中之一，即选取s[i][j-1]和s[i+1][j]的最大值
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)    
	for i := range dp {
		dp[i] = make([]int, n)
		// 单个字母也是长度为1的回文字符串
		dp[i][i] = 1
	}

	// 从下向上填充一个上三角矩阵，返回矩阵右上角
	for i:=n-2; i>=0; i-- {
		for j:=i+1; j<n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			}else{
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][n-1]
}

func main() {
	/* bbbab的dp输出
	[1 2 3 3 4] 
	[0 1 2 2 3] 
	[0 0 1 1 3] 
	[0 0 0 1 1] 
	[0 0 0 0 1]
	*/
	s := []string{"bbbab", "cbbd"}
	for i := range s {
		fmt.Println(longestPalindromeSubseq(s[i]))
	}
}
