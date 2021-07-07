package main

import (
	"fmt"
)

// 在0078_Subsets/naive.go的基础上在for内部增加一个if判断条件去重
// 类似的问题：0039_Combination_Sum和0040_Combination_Sum_II
func backtrack(nums []int, rslt *[][]int, curr []int, currIdx int) {
	valid := make([]int, len(curr), cap(curr))
	copy(valid, curr)
	*rslt = append(*rslt, valid)	

	for i:=currIdx; i<len(nums); i++ {
		if i > currIdx && nums[i-1] == nums[i] {
			continue
		}
		curr = append(curr, nums[i])
		backtrack(nums, rslt, curr, i+1)
		curr = curr[0:len(curr)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	var rslt [][]int
	backtrack(nums, &rslt, []int{}, 0)
	return rslt 
}

func main() {
	fmt.Println(subsetsWithDup([]int{1,2,2}))
	fmt.Println(subsetsWithDup([]int{0}))
}
