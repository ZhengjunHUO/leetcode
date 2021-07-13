package main

import (
	"fmt"
)

/*
m*n为3*7的情况: 每个格子的值代表从它出发可以选择的路径条数
_ 21 15 10 6 3 1 
_  6  5  4 3 2 1
_  1  1  1 1 1 1
保持一个长为m的存储，从右向左计算并覆盖，最后对其求和(21+6+1=28)
*/

func uniquePaths(m int, n int) int {
	if n == 1 || m == 1 {
		return 1
	}

	if n == 2 {
		return m
	}

	if m == 2 {
		return n
	}

	base := make([]int, m) 
	for k := range base {
		base[k] = 1
	}

	rslt := 0	

	for i:=0; i<n-2; i++ {
		for j:=1; j<m; j++ {
			base[j] = base[j-1] + base[j]
		}	
	}

	for _,v := range base {
		rslt += v
	}

	return rslt
}

func main() {
	m := []int{3, 3, 7, 3}
	n := []int{7, 2, 3, 3}

	for i := range m {
		fmt.Println(uniquePaths(m[i], n[i]))
	}
}
