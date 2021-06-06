package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	n, cnt := len(nums), 0 

	for i:=0; i<n; i++ {
		sum := 0
		for j:=i; j<n; j++ {
			sum += nums[j]
			if sum == k {
				cnt += 1
			}
		}
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
}
