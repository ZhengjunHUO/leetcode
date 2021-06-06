package main

import (
	"fmt"
)

func checkSubarraySum(nums []int, k int) bool {
	cumsum := make(map[int]int)
	cumsum[0] = -1
	sum := 0

	for i := range nums {
		sum += nums[i]
		sum %= k

		if v,ok := cumsum[sum]; ok {
			if i - v > 1 {
				return true
			}	
		}	
		cumsum[sum]=i
	}
	return false
}

func main() {
	nums := []int{23,2,4,6,7}
	k := 6
	fmt.Println(checkSubarraySum(nums, k))

	nums = []int{23,2,6,4,7}
	k = 13
	fmt.Println(checkSubarraySum(nums, k))
}
