package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func binaryTreePaths(root *godtype.TreeNode) []string {
	rslt := []string{}
	var recursive func(*godtype.TreeNode, string)

	recursive = func(curr *godtype.TreeNode, currPath string) {
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
		fmt.Println(binaryTreePaths(godtype.NewBTree(trees[i])))
	}
}
