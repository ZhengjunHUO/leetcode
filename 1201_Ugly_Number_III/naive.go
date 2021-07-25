package main

import (
	"fmt"
)

// 修改自0264_Ugly_Number_II/solution.go
func min(a, b, c int) int {
	var temp int
	if a < b {
		temp = a
	}else{
		temp = b
	}

	if temp > c {
		temp = c
	}

	return temp
}

func nthUglyNumber(n int, a int, b int, c int) int {
	list := make([]int, n)

	ia, ib, ic := 1, 1, 1
	for i:=0; i<n; i++ {
		next := min(ia*a, ib*b, ic*c)
		list[i] = next

		if next == ia*a {
			ia++
		}
		if next == ib*b {
			ib++ 
		}
		if next == ic*c {
			ic++
		}
	}

	return list[n-1]
}

func main() {
	nums := [][]int{[]int{3, 2, 3, 5}, []int{4, 2, 3, 4}, []int{5, 2, 11, 13}, []int{1000000000, 2, 217983653, 336916467}}
	for _, v := range nums {
		fmt.Println(nthUglyNumber(v[0], v[1], v[2], v[3]))
	}
}
