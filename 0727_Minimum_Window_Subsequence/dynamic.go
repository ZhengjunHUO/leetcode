package main

import (
	"fmt"
)

func minWindow(s,t string) string {
	len_s, len_t := len(s), len(t)

	/*
	以目标字符串t（较短）为第一维度（行），逐行填充矩阵
	矩阵中非零内容为"t字符串中第一个字符"在s字符串中出现的index
	最后一行（此时已比对了所有t中的字符）中的uniq(非零值)为最终解
	本题中还需要比较长度，返回第一个最短的子字符串
	*/
	dp := make([][]int, len_t+1)
	for x := range dp {
		dp[x] = make([]int, len_s+1)
	}

	start_idx, minLen := -1, 20000

	// 初始化第0行，便于第1行在匹配的情况下获取index
	for i:=0; i<len_s+1; i++ {
		dp[0][i] = i
	}

	// 逐行填充矩阵，若匹配则继承左上格的值；不匹配则继承左边的值
	for i:=1; i<len_t+1; i++ {
		for j:=1; j<len_s+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	// 遍历最后一行，找到最优解
	for i:=0; i<len_s+1; i++ {
		if begin_idx := dp[len_t][i]; begin_idx != 0 {
			temp := i - begin_idx 
			if temp < minLen {
				start_idx = begin_idx 
				minLen = temp
			}
		}
	}

	if minLen < 20000 {
		return s[start_idx:start_idx+minLen]
	}

	return ""
}

func main() {
	s, t := "abcdebdde", "bde"
	fmt.Println(minWindow(s,t))
}


/* 
dp中的值
    a b c d e b d d e
  0 1 2 3 4 5 6 7 8 9
b 0 0 1 1 1 1 5 5 5 5
d 0 0 0 0 1 1 1 5 5 5
e 0 0 0 0 0 1 1 1 1 5
*/

