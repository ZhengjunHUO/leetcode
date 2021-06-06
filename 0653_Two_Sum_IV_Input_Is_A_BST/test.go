package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		5,
		&TreeNode{
			3,
//			&TreeNode{},
//			&TreeNode{},
			nil,
			nil,
		},
		&TreeNode{
			6,
			&TreeNode{},
			&TreeNode{7, &TreeNode{}, &TreeNode{}},
		},
	}

	if root.Left.Left == nil {
		fmt.Println("is nil")
	}
//	fmt.Println(root.Left.Left)
}
