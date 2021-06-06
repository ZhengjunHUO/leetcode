package main

import (
	"fmt"
)

func subarraysDivByK(nums []int, k int) int {
	cumsum := make(map[int]int)
	cumsum[0] = 1

	sum,cnt := 0,0

	for i := range nums {
		sum += nums[i]

		r := sum%k
		if r < 0 {
			r += k
		}

		if v,ok := cumsum[r]; ok {
			cnt += v
		}
		cumsum[r] += 1
	}
	
	return cnt
    
}

func main() {
	nums := []int{4,5,0,-2,-3,1}
	k := 5
	fmt.Println(subarraysDivByK(nums,k))

	nums = []int{3,4,7,2,-3,1,4,2}
	k = 7
	fmt.Println(subarraysDivByK(nums,k))
}
