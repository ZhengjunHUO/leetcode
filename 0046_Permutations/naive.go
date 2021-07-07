package main

import (
	"fmt"
)

func backtrack(nums []int, rslt *[][]int, curr []int, dict map[int]int) {
	if len(curr) == len(nums) {
		valid := make([]int, len(nums), cap(nums))
		copy(valid, curr)
		*rslt = append(*rslt, valid)
		return
	}

	for i:=0; i<len(nums); i++ {
		if _, ok := dict[nums[i]]; ok {
			continue
		}
		dict[nums[i]] = 1
		curr = append(curr, nums[i])

		backtrack(nums, rslt, curr, dict)

		delete(dict, nums[i])
		curr = curr[0:len(curr)-1]
	}
}

func permute(nums []int) [][]int {
	var rslt [][]int
	dict := make(map[int]int)
	backtrack(nums, &rslt, []int{}, dict)
	return rslt
}

func main() {
	fmt.Println(permute([]int{1,2,3}))
	fmt.Println(permute([]int{0,1}))
	fmt.Println(permute([]int{1}))
}
