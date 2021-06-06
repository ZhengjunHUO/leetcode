package main

import (
	"fmt"
)

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, n+1)
	for i:=0; i<=n; i++ {
		dp[i] = make([]int, m+1)
	}

	for _,s := range strs {
		ones, zeros := 0, 0
		for _,i := range s {
			if i == '1' {
				ones += 1
			}
			if i == '0' {
				zeros += 1
			}
		}

		for i:=n; i>=ones; i-- {
			for j:=m; j>=zeros; j-- {
				if temp := dp[i-ones][j-zeros] + 1; temp > dp[i][j] {
					dp[i][j] = temp
				}
			}
		}	
	}
	
	return dp[n][m]
}

func main() {
	strs := []string{"10","0001","111001","1","0"}
	m, n := 5, 3
	fmt.Println(findMaxForm(strs, m, n))

	strs = []string{"10","0","1"}
	m, n = 1, 1
	fmt.Println(findMaxForm(strs, m, n))
}
