package main

import (
	"github.com/ZhengjunHUO/godtype"	
)

func max(nums []int) int {
	max, index := -1, -1
	for i := range nums {
		if nums[i] > max {
			max, index = nums[i], i
		}
	}	

	return index
}


func constructMaximumBinaryTree(nums []int) *godtype.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	idx := max(nums)
	return &godtype.TreeNode{
		Val: nums[idx],
		Left: constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
        }
}

func main() {
	nums := [][]int{[]int{3,2,1,6,0,5}, []int{3,2,1}}
	for i := range nums {
		godtype.PrintBTreeDFS(constructMaximumBinaryTree(nums[i]))
	}
}
