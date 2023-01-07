package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
)

func searchBST(root *graph.TreeNode, val int) *graph.TreeNode {
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
	tree := graph.NewBTree([]interface{}{4,2,7,1,3})
	vals := []int{2, 5}
	for _, v := range vals {
		graph.PrintBTreeBFS(os.Stdout, searchBST(tree, v))
	}
}
