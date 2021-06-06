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
	dict := make(map[int]int)    
	return scanChildren(root, k, dict)
}

func scanChildren(current *TreeNode, target int, dict map[int]int) bool {
	if current == nil {
		return false
	}

	diff := target - current.Val	
	if _, ok := dict[diff]; ok {
		return true
	}else{
		dict[current.Val] = 1
		return scanChildren(current.Left, target, dict) || scanChildren(current.Right, target, dict)
	}
}


func main() {
	root := &TreeNode{
		5,
		&TreeNode{
			3,
			&TreeNode{2, nil, nil}, 
			&TreeNode{4, nil, nil}, 
		},
		&TreeNode{
			6,
			nil,
			&TreeNode{7, nil, nil}, 
		},
	}

	fmt.Println(findTarget(root, 9))
	fmt.Println(findTarget(root, 28))
}
