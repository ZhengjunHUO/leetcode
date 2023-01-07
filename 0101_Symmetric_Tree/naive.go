package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func recurs(left, right *graph.TreeNode) bool {
	if ( left == nil || right == nil) {
		return left == right
	}

	if left.Val.(int) != right.Val.(int) {
		return false
	}

	return recurs(left.Left, right.Right) && recurs(left.Right, right.Left)
}

func isSymmetric(root *graph.TreeNode) bool {
	return root != nil && recurs(root.Left, root.Right)
}

func main() {
	trees := [][]interface{}{[]interface{}{1,2,2,3,4,4,3}, []interface{}{1,2,2,nil,3,nil,3}}
	for i := range trees {
		fmt.Println(isSymmetric(graph.NewBTree(trees[i])));
	}
}
