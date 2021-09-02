package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	n := len(nums)
	rslt := n
	for i := range nums {
		rslt ^= i ^ nums[i]
	}

	return rslt
}

func main() {
	nums := [][]int{[]int{3,0,1}, []int{0,1}, []int{9,6,4,2,3,5,7,0,1}, []int{0}}
	for i := range nums {
		fmt.Println(missingNumber(nums[i]))
	}
}
