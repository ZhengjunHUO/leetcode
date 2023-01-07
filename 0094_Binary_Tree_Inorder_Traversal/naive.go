package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func inorderTraversal(root *graph.TreeNode) []int {
	rslt := []int{}
	var recursive func(*graph.TreeNode)

	recursive = func(curr *graph.TreeNode) {
		if curr == nil {
			return
		}

		recursive(curr.Left)
		rslt = append(rslt, curr.Val.(int))
		recursive(curr.Right)
	}

	recursive(root)
	return rslt
}

func main() {
	ts := [][]interface{}{[]interface{}{1,nil,2,nil,nil,3}, []interface{}{}, []interface{}{1}, []interface{}{1,2}, []interface{}{1,nil,2}}
	for i := range ts {
		fmt.Println(inorderTraversal(graph.NewBTree(ts[i])))
	}
}
