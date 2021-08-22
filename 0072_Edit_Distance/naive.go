package main

import (
	"fmt"
)

func min3(a, b, c int) int {
	var temp int
	if a < b {
		temp = a
	}else{
		temp = b
	}

	if temp < c {
		return temp
	}

	return c
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1) + 1)
	for i := range dp {
		dp[i] = make([]int, len(word2) + 1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 1; i < len(word1) + 1; i++ {
		for j := 1; j < len(word2) + 1; j++ {
			// 如word1和word2当前字符相等，不做操作跳过
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else{
				// 取插入，删除，替换操作的最小操作数
				dp[i][j] = min3(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func main() {
	words1 := []string{"horse", "intention"}
	words2 := []string{"ros", "execution"}

	for i := range words1 {
		fmt.Println(minDistance(words1[i], words2[i]))
	}
}
