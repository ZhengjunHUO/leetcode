package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
)

func invertTree(root *graph.TreeNode) *graph.TreeNode {
	if root == nil {
		return nil
	}

	invert(root)
	return root
}

func invert(curr *graph.TreeNode) {
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
		graph.PrintBTreeBFS(os.Stdout, invertTree(graph.NewBTree(trees[i])))
	}
}
