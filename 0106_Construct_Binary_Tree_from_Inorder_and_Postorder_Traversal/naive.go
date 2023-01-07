package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
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

func buildTree(inorder []int, postorder []int) *graph.TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	// postorder的最后一个元素为当前树的根，通过查找inorder来确定左子树的长度
	// postorder中右子树的范围是头部去掉左子树的长度，尾部去掉长度为1的根的中间部分
	idx := findValue(inorder, postorder[len(postorder)-1])
	return &graph.TreeNode{
		Val: postorder[len(postorder)-1],
		Left: buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
        }
}

func main() {
	graph.PrintBTreeDFS(os.Stdout, buildTree([]int{9,3,15,20,7}, []int{9,15,7,20,3}))
	graph.PrintBTreeDFS(os.Stdout, buildTree([]int{-1}, []int{-1}))
}
