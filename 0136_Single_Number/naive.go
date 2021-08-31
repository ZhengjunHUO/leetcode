package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	rslt := 0
	for i := range nums {
		rslt ^= nums[i]
	}

	return rslt
}

func main() {
	nums := [][]int{[]int{2,2,1}, []int{4,1,2,1,2}, []int{1}}
	for i := range nums {
		fmt.Println(singleNumber(nums[i]))
	}
}
