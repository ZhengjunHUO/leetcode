package main

import (
	"fmt"
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
)

const max_val int = (1 << 31)
const min_val int = -(1 << 31)-1

func recover(curr, minNode, maxNode *graph.TreeNode) {
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

func recoverTree(root *graph.TreeNode)  {
	recover(root, nil, nil)
}

func main() {
	trees := [][]interface{}{[]interface{}{1,3,nil,nil,2}, []interface{}{3,1,4,nil,nil,2}}
	for i := range trees {
		t := graph.NewBTree(trees[i])
		fmt.Printf("Before: ")
		graph.PrintBTreeBFS(os.Stdout, t)

		recoverTree(t)
		fmt.Printf("After : ")
		graph.PrintBTreeBFS(os.Stdout, t)
	}
}
