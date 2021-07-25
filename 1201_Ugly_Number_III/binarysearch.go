package main

import (
	"fmt"
)

// 参考@hiepit的算法
// 求最大公约数
func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b % a, a)
}

// 求最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 计算小于m的ugly number个数
func calculate(m, a, b, c int) int{
	return m/a + m/b + m/c - m/lcm(a,b) - m/lcm(a,c) - m/lcm(b,c) + m/lcm(a, lcm(b,c))
}

// 使用二分法
func nthUglyNumber(n int, a int, b int, c int) int {
	rslt, l ,r := 0, 0, int(2e9)

	for l <= r {
		m := l + ( r - l )/2
		if temp := calculate(m, a, b, c); temp >= n {
//			fmt.Println("calculate: ", temp, "; result updated: ", m)
			rslt = m
			r = m - 1
		}else{
//			fmt.Println("calculate: ", temp)
			l = m + 1
		}
	}

	return rslt
}

func main() {
	nums := [][]int{[]int{3, 2, 3, 5}, []int{4, 2, 3, 4}, []int{5, 2, 11, 13}, []int{1000000000, 2, 217983653, 336916467}}
	for _, v := range nums {
		fmt.Println(nthUglyNumber(v[0], v[1], v[2], v[3]))
	}
}
