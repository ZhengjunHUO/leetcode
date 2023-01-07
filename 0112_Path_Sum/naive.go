package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func backtrack(current *graph.TreeNode, remain int, hasPath *bool) {
	if current == nil {
		if remain == 0 {
			*hasPath = true
		}
		return
	}

	if *hasPath == true {
		return
	}

	curVal := current.Val.(int)
	backtrack(current.Left, remain - curVal, hasPath)
	backtrack(current.Right, remain - curVal, hasPath)
}

func hasPathSum(root *graph.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var hasPath bool = false
	backtrack(root, targetSum, &hasPath)

	return hasPath
}

func main() {
	trees := [][]interface{}{[]interface{}{5,4,8,11,nil,13,4,7,2,nil,nil,nil,1}, []interface{}{1,2,3}, []interface{}{}}
	targetSums := []int{22,5,0}
	for i := range trees {
		fmt.Println(hasPathSum(graph.NewBTree(trees[i]), targetSums[i]))
	}
}
