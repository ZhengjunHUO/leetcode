package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func binaryTreePaths(root *graph.TreeNode) []string {
	rslt := []string{}
	var recursive func(*graph.TreeNode, string)

	recursive = func(curr *graph.TreeNode, currPath string) {
		currPath = currPath + fmt.Sprint(curr.Val.(int)) + "->"

		if curr.Left == nil && curr.Right == nil {
			rslt = append(rslt, currPath[:len(currPath)-2])
			return
		}

		if curr.Left != nil {
			recursive(curr.Left, currPath)
		}
		if curr.Right != nil {
			recursive(curr.Right, currPath)
		}
	}

	recursive(root, "")
	return rslt
}

func main() {
	trees := [][]interface{}{[]interface{}{1,2,3,nil,5}, []interface{}{1}}
	for i := range trees {
		fmt.Println(binaryTreePaths(graph.NewBTree(trees[i])))
	}
}
