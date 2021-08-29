package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func build(nums []int, l, r int) *godtype.TreeNode {
	if l > r {
		return nil
	}

	mid := l + (r-l) / 2
	
	return &godtype.TreeNode{
		Val:	nums[mid],
		Left:	build(nums, l, mid-1),
		Right:	build(nums, mid+1, r),
	}
}

func sortedArrayToBST(nums []int) *godtype.TreeNode {
	return build(nums, 0, len(nums)-1)	
}

func main() {
	nums := [][]int{[]int{-10,-3,0,5,9}, []int{1, 3}}
	for i := range nums {
		godtype.PrintBTreeBFS(sortedArrayToBST(nums[i]))
	}
}
