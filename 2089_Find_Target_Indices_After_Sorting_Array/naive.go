package main

import (
	"fmt"
	"sort"
)

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	idx := sort.SearchInts(nums, target)

	ret := []int{}
	if nums[idx] != target {
		return ret		
	}

	for idx < len(nums) && nums[idx] == target {
		ret = append(ret, idx)
		idx++
	}

	return ret
}

func main() {
	nums := []int{1,2,5,2,3}
	targets := []int{2,3,5,4}
	for i := range targets {
		fmt.Println(targetIndices(nums, targets[i]))
	}
}
