package main

import (
	"fmt"
	"sort"
)

func twoSumLessK(nums []int, target int) int {
	sort.Ints(nums)
	lp, rp, rslt := 0, len(nums)-1, -1
	for (lp<rp) {
		if sum := nums[lp]+nums[rp]; sum > target {
			rp -= 1
		}else{
			if sum > rslt {
				rslt = sum
			}
			lp += 1	
		}
	}
	return rslt
}

func main() {
	nums := []int{34,23,1,24,75,33,54,8}
	target := 50
	fmt.Println(twoSumLessK(nums, target))
}
