package main

import (
	"fmt"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	cnt := 0
	lp, rp, sum := 0, 0, nums[0]
	for ! (lp==rp && rp==(len(nums)-1)) {
		if sum < k {
			cnt += 1 
			fmt.Println(lp,rp)
		}

		if sum < k && rp < len(nums)-1 {
			rp += 1
			sum *= nums[rp]
			continue
		}

		sum /= nums[lp]
		lp += 1
		if nums[lp] < k {
			cnt += 1
		}
	}
   	return cnt 
}

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
	fmt.Println(numSubarrayProductLessThanK(nums, k))
}
