package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

const max_val int = (1 << 31)
const min_val int = -(1 << 31)-1

func recover(curr, minNode, maxNode *godtype.TreeNode) {
	if curr == nil {
		return
	}

	if minNode != nil && curr.Val.(int) < minNode.Val.(int) {
		minNode.Val, curr.Val = curr.Val, minNode.Val
		return
	}

	if maxNode != nil && curr.Val.(int) > maxNode.Val.(int) {
		maxNode.Val, curr.Val = curr.Val, maxNode.Val
		return
	}

	recover(curr.Left, minNode, curr)
	recover(curr.Right, curr, maxNode)
}

func recoverTree(root *godtype.TreeNode)  {
	recover(root, nil, nil)
}

func main() {
	trees := [][]interface{}{[]interface{}{1,3,nil,nil,2}, []interface{}{3,1,4,nil,nil,2}}
	for i := range trees {
		t := godtype.NewBTree(trees[i])
		fmt.Printf("Before: ")
		godtype.PrintBTreeBFS(t)

		recoverTree(t)
		fmt.Printf("After : ")
		godtype.PrintBTreeBFS(t)
	}
}
