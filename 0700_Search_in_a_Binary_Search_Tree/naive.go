package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func searchBST(root *godtype.TreeNode, val int) *godtype.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val.(int) > val {
		return searchBST(root.Left, val)
	}
	if root.Val.(int) < val {
		return searchBST(root.Right, val)
	}

	return root
}

func main() {
	tree := godtype.NewBTree([]interface{}{4,2,7,1,3})
	vals := []int{2, 5}
	for _, v := range vals {
		godtype.PrintBTreeBFS(searchBST(tree, v))
	}
}
