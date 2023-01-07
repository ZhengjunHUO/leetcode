package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func postorderTraversal(root *graph.TreeNode) []int {
	rslt := []int{}
	var recursive func(*graph.TreeNode)

	recursive = func(curr *graph.TreeNode) {
		if curr == nil {
			return
		}

		recursive(curr.Left)
		recursive(curr.Right)
		rslt = append(rslt, curr.Val.(int))
	}

	recursive(root)
	return rslt
}

func main() {
	ts := [][]interface{}{[]interface{}{1,nil,2,nil,nil,3}, []interface{}{}, []interface{}{1}, []interface{}{1,2}, []interface{}{1,nil,2}}
	for i := range ts {
		fmt.Println(postorderTraversal(graph.NewBTree(ts[i])))
	}
}
