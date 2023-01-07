package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func search(curr *graph.TreeNode, min, max int) int {
	if curr == nil {
		return 100001
	}

	val, rslt := curr.Val.(int), 100001

	if min != -1 {
		rslt = min2(rslt, val-min)
	}
	if max != 100001 {
		rslt = min2(rslt, max-val)
	}

	rslt = min2(rslt, search(curr.Left, min, val))
	rslt = min2(rslt, search(curr.Right, val, max))

	return rslt
}

func getMinimumDifference(root *graph.TreeNode) int {
	return search(root, -1, 100001)
}

func main() {
	trees := [][]interface{}{[]interface{}{4,2,6,1,3}, []interface{}{1,0,48,nil,nil,12,49}, []interface{}{12,1,36,nil,nil,14,49}}
	for i := range trees {
		fmt.Println(getMinimumDifference(graph.NewBTree(trees[i])))
	}
}
