package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, target, 4)
}

// nums should be a sorted list
func nSum(nums []int, target, n int) [][]int {
	if len(nums)==0 || nums[0]*n > target || nums[len(nums)-1] * n < target {
		return [][]int{}
	}

	rslt := [][]int{}
	if n==2 {
		return twoSum(nums, target)
	}

	for i := range nums {
		if i==0 || nums[i-1]!=nums[i] {
			for _,elem := range nSum(nums[i+1:], target-nums[i], n-1) {
				temp := []int{nums[i]}
				temp = append(temp, elem...)
				rslt = append(rslt, temp)
			}
		}
	}
	return rslt
}

func twoSum(nums []int, target int) [][]int {
	rslt := [][]int{}
	lp, rp := 0, len(nums)-1
	for (lp < rp) {
		sum:=nums[lp]+nums[rp]
		switch {
		case sum==target:
			rslt = append(rslt, []int{nums[lp],nums[rp]})
			lp += 1
			rp -= 1
		case sum<target || (lp>0 && nums[lp]==nums[lp-1]):
			lp += 1
		case sum>target || (rp<len(nums)-1 && nums[rp]==nums[rp+1]):
			rp -= 1
		}
	}
	return rslt
}

func main() {
	nums := []int{1,0,-1,0,-2,2}
	target := 0
	fmt.Println(fourSum(nums, target))
}
