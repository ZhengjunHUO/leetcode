package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func minDepth(root *graph.TreeNode) int {
	if root == nil {
		return 0
	}

	q := datastruct.NewQueue([]*graph.TreeNode{})
	q.Push(root)
	depth := 0

	for !q.IsEmpty() {
		depth++

		size := q.Size()
		for i:=0; i<size; i++ {
			node := q.Pop()

			// 遇到叶子结点，返回
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}

	return depth
}

func main() {
	fmt.Println(minDepth(graph.NewBTree([]interface{}{3,9,20,nil,nil,15,7})))
}
