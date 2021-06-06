package main

import (
	"fmt"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	cnt,lp,prod := 0,0,1

	for rp := range nums {
		prod *= nums[rp]
		for prod >= k {
			prod /= nums[lp]
			lp += 1
		}
		cnt += rp - lp + 1
	}

   	return cnt 
}

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
	fmt.Println(numSubarrayProductLessThanK(nums, k))
}
