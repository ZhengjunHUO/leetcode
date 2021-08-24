package main

import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	dp, prev := make([]int, len(text2)+1), make([]int, len(text2)+1)

	for i := range text1 {
		for j := 1; j < len(dp); j++ {
			// 如果字符相等，取左上角的值加1
			// 即dp[i][j] = dp[i-1][j-1] + 1
			if text1[i] == text2[j-1] {
				dp[j] = prev[j-1]+1
				continue
			} 
			// 如不等，则取较大历史记录的一边
			// 即dp[i][j] == max(dp[i][j-1], dp[i-1][j])
			if dp[j-1] > prev[j] {
				dp[j] = dp[j-1]
			}
		}
		copy(prev, dp)
	}

	return dp[len(dp)-1]
}

func main() {
	s1 := []string{"abcde", "abc", "abc"}
	s2 := []string{"ace", "abc", "def"}

	for i := range s1 {
		fmt.Println(longestCommonSubsequence(s1[i], s2[i]))
	}
}
