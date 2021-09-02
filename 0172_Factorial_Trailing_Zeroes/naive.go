package main

import (
	"fmt"
)

/*
func count2(n int) int {
	temp := n/2
	count := 0
	for n > 0 {
		n = n >> 1
		count++
	}

	if count <= 2 {
		return temp
	}
	return (1 + (count-2))*(count-2)/2 + temp
}
*/

func count5(n int) int {
	temp := n/5
	count := 0
	for n > 0 {
		n /= 5
		count++
	}

	if count <= 2 {
		return temp
	}
	return (1 + (count-2))*(count-2)/2 + temp
}

/* 
  结尾0的个数取决于小于等于n的数中2和5因子的个数
  因子2的个数永远大于因子5的个数，所以只需统计5的个数
*/
func trailingZeroes(n int) int {
	/*
	c2, c5 := count2(n), count5(n)
	if c2 > c5 {
		return c5
	}
	return c2
	*/

	return count5(n)
}

func main() {
	ns := []int{3,5,0,20,25}
	// 20: 4个0; 25: 6个0
	for i := range ns {
		fmt.Println(trailingZeroes(ns[i]))
	}
}
