package main

import (
	"fmt"
)

/*
m*n的矩阵总共需要向右走n-1步，向下走m-1步
eg. 下下右下下下右下下
即求在长度为(m-1)+(n-1)的路线图中选择m个位置(或n个位置)的可能性种类
C(m, m+n) => (m+n)! / ((m+n - m)! * m!) => (m+n)!/(n!)/(m!) => (m+1)*(m+2)*...*(m+n)/n!
*/
func uniquePaths(m int, n int) int {
	if n == 1 || m == 1 {
		return 1
	}

	m -= 1
	n -= 1

	// 确保m>n
	if m < n {
		m, n = n, m
	}

	var rslt float64 = 1
	var div float64 = 1
	// 实现(m+1)*(m+2)*...*(m+n)/n!
	for i:=m+1; i<=m+n; i++ {
		rslt = rslt*float64(i)/div
		div += 1
	}

	return int(rslt)
}

func main() {
	m := []int{3, 3, 7, 3}
	n := []int{7, 2, 3, 3}

	for i := range m {
		fmt.Println(uniquePaths(m[i], n[i]))
	}
}
