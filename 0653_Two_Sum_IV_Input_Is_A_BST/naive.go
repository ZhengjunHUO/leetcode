package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func findTarget(root *TreeNode, k int) bool {
    
}


func main() {
	root := &TreeNode{
		5,
		&TreeNode{
			3,
			&TreeNode{2, &TreeNode{}, &TreeNode{}},
			&TreeNode{4, &TreeNode{}, &TreeNode{}},
		},
		&TreeNode{
			6,
			&TreeNode{},
			&TreeNode{7, &TreeNode{}, &TreeNode{}},
		},
	}

	fmt.Println(findTarget(root, 9))
	fmt.Println(findTarget(root, 28))
}
