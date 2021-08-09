package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func invertTree(root *godtype.TreeNode) *godtype.TreeNode {
	if root == nil {
		return nil
	}

	invert(root)
	return root
}

func invert(curr *godtype.TreeNode) {
	if curr.Left == nil && curr.Right == nil {
		return 
	}

	curr.Left, curr.Right = curr.Right, curr.Left
	if curr.Left != nil {
		invert(curr.Left)
	}
	if curr.Right != nil {
		invert(curr.Right)
	}
	return
}

func main() {
	trees := [][]interface{}{[]interface{}{4,2,7,1,3,6,9}, []interface{}{2,1,3}, []interface{}{}}
	for i := range trees {
		godtype.PrintBtreeBFS(invertTree(godtype.NewBTree(trees[i],0)))
	}
	
}
