package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

var maxDist int

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func search(curr *godtype.TreeNode) int {
	if curr == nil {
		return 0
	}

	l, r := search(curr.Left), search(curr.Right)
	if l + r > maxDist {
		maxDist = l + r
	}

	return max(l, r) + 1
}

func diameterOfBinaryTree(root *godtype.TreeNode) int {
	maxDist = 0
	search(root)
	return maxDist
}

func main() {
	trees := [][]interface{}{[]interface{}{1,2,3,4,5}, []interface{}{1,2}}
	for i := range trees {
		fmt.Println(diameterOfBinaryTree(godtype.NewBTree(trees[i])))
	}
}
