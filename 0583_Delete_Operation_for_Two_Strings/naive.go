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

// 类似题目0072_Edit_Distance
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// 初值为0-n，即需要等于空字符串需要执行的删除操作数(递增加1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
		
	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			if word1[i-1] == word2[j-1] {
				// 如当前字符相等，直接继承左上角的值
				dp[i][j] = dp[i-1][j-1]	
			}else{
				// dp[i][j-1](左)为word1为了和word2相等需要的删除操作数, 另一部分同理
				// 取其最小值并加上当前操作
				dp[i][j] = min2(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}

func main() {
	words1 := []string{"sea", "leetcode", "international"}
	words2 := []string{"eat", "etco", "lions"}

	for i := range words1 {
		fmt.Println(minDistance(words1[i], words2[i]))
	}
}
