package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
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


func constructMaximumBinaryTree(nums []int) *graph.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	idx := max(nums)
	return &graph.TreeNode{
		Val: nums[idx],
		Left: constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
        }
}

func main() {
	nums := [][]int{[]int{3,2,1,6,0,5}, []int{3,2,1}}
	for i := range nums {
		graph.PrintBTreeDFS(os.Stdout, constructMaximumBinaryTree(nums[i]))
	}
}
