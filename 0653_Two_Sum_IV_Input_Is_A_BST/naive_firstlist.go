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
	list := make([]int,0,4)
//	fmt.Printf("[Init] Pointer's addr: %p\n", &list)
	parseTree(root, &list)    
//	fmt.Println("Get list: ", list)

	lp, rp := 0, len(list) - 1
	for (lp < rp) {
		sum := list[lp] + list[rp] 
		switch {
		case sum == k:
			return true
		case sum < k:
			lp += 1
		case sum > k:
			rp -= 1
		}
	}
	return false
}

func parseTree(current *TreeNode, list *[]int) {
	if current.Left != nil {
		parseTree(current.Left, list)
	}
	*list = append(*list, current.Val)
//	fmt.Printf("Pointer's addr: %p\n", list)
//	fmt.Printf("Pointer's value: %p\n", *list)
//	fmt.Printf("First elem's addr: %p\n", &(*list)[0])
	if current.Right != nil {
		parseTree(current.Right, list)
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
