package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	cumsum := make(map[int]int)
	cumsum[0] = 1
	
	cnt, sum := 0, 0

	for i := range nums {
		sum += nums[i]
		if v,ok := cumsum[sum-k]; ok {
			cnt += v
		}
		cumsum[sum]+=1
	}

	return cnt
}

func main() {
	nums := []int{1,1,1}
	k := 2
	fmt.Println(subarraySum(nums, k))

	nums = []int{1,2,3}
	k = 3
	fmt.Println(subarraySum(nums, k))

	nums = []int{3,4,7,2,-3,1,4,2}
	k = 7
	fmt.Println(subarraySum(nums, k))

	nums = []int{3,4,7,2,0,-3,1,4,2}
	k = 0
	fmt.Println(subarraySum(nums, k))
}
