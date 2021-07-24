package main

import (
	"fmt"
)

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

func nthUglyNumber(n int) int {
	list := make([]int, n)
	list[0] = 1

	i2, i3, i5 := 0, 0, 0
	for i:=1; i<n; i++ {
		next := min(list[i2]*2, list[i3]*3, list[i5]*5)
		list[i] = next

		if next == list[i2]*2 {
			i2++
		}
		if next == list[i3]*3 {
			i3++ 
		}
		if next == list[i5]*5 {
			i5++
		}
	}

	return list[n-1]
}

func main() {
	nums := []int{10, 1, 20}
	for _, v := range nums {
		fmt.Println(nthUglyNumber(v))
	}
}
