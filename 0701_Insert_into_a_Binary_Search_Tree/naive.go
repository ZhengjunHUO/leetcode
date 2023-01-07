package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
)

func insertIntoBST(root *graph.TreeNode, val int) *graph.TreeNode {
	if root == nil {
		return &graph.TreeNode{val, nil, nil}
	}

	if val < root.Val.(int) {
		root.Left = insertIntoBST(root.Left, val)
	}else{
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

func main() {
	lists := [][]interface{}{[]interface{}{4,2,7,1,3}, []interface{}{40,20,60,10,30,50,70}}
	vals := []int{5,25}
	for i := range lists {
		graph.PrintBTreeBFS(os.Stdout, insertIntoBST(graph.NewBTree(lists[i]), vals[i]))
	}

}
