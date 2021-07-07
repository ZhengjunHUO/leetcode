package main

import (
	"fmt"
)

func backtrack(nums []int, rslt *[][]int, curr []int, currIdx int) {
	valid := make([]int, len(curr), cap(curr))
	copy(valid, curr)
	*rslt = append(*rslt, valid)	

	/*
	if currIdx == len(nums) {
		return
	}
	*/

	for i:=currIdx; i<len(nums); i++ {
		curr = append(curr, nums[i])
		backtrack(nums, rslt, curr, i+1)
		curr = curr[0:len(curr)-1]
	}
}

func subsets(nums []int) [][]int {
	var rslt [][]int
	backtrack(nums, &rslt, []int{}, 0)
	return rslt 
}

func main() {
	fmt.Println(subsets([]int{1,2,3}))
	fmt.Println(subsets([]int{0}))
}
