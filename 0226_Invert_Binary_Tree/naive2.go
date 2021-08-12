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
	if curr == nil {
		return 
	}

	curr.Left, curr.Right = curr.Right, curr.Left
	invert(curr.Left)
	invert(curr.Right)
}

func main() {
	trees := [][]interface{}{[]interface{}{4,2,7,1,3,6,9}, []interface{}{2,1,3}, []interface{}{}}
	for i := range trees {
		godtype.PrintBTreeBFS(invertTree(godtype.NewBTree(trees[i])))
	}
	
}
