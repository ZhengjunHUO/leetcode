package main

import (
	"fmt"
)

func calc(n int) int {
	return n * (n+1) / 2
}

// 二分搜索行数
func arrangeCoins(n int) int {
	l, r := 1, n
	for l < r {
		m := l + (r - l)/2
		// 正好填满，返回m
		if calc(m) == n {
			return m
		}

		// 搜索左边界
		if calc(m) > n {
			r = m
		}else{
			l = m + 1
		}
	}

	// 根据题目要求未填满的行不算入结果，需要减1（填满的情况以及直接返回）
	return l - 1
}

func main() {
	ns := []int{5,8,6}
	for i := range ns {
		fmt.Println(arrangeCoins(ns[i]))
	}
}
