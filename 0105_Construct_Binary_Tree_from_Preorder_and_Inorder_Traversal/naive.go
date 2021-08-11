package main

import (
	"github.com/ZhengjunHUO/godtype"
)

// 找到元素在slice中的index
func findValue(nums []int, target int) int {
	for i := range nums {
		if target == nums[i] {
			return i
		}
	}

	return -1
}

func buildTree(preorder []int, inorder []int) *godtype.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// preorder的第一个元素为当前树的根，通过查找inorder来确定左子树的长度
	// preorder和inorder去掉第一个元素和左子树的长度剩余部分为右子树
	idx := findValue(inorder, preorder[0])
	return &godtype.TreeNode{
		Val: preorder[0], 
		Left: buildTree(preorder[1:1+idx], inorder[:idx]),
		Right: buildTree(preorder[1+idx:], inorder[idx+1:]),
        }
}

func main() {
	godtype.PrintBtree(buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7}))
	godtype.PrintBtree(buildTree([]int{-1}, []int{-1}))
}
