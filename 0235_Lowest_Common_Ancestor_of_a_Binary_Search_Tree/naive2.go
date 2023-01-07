package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func lowestCommonAncestor(root *graph.TreeNode, p, q int) *graph.TreeNode {
	if p > q {
		p, q = q, p
	}

	for ! (root.Val.(int) >= p && root.Val.(int) <= q) {
		if root.Val.(int) < p {
			root = root.Right
		}else{
			root = root.Left
		}
	}

	return root
}

func main() {
	ps := []int{2,2,0}
	qs := []int{8,4,5}

	for i := range ps {
		fmt.Println(lowestCommonAncestor(graph.NewBTree([]interface{}{6,2,8,0,4,7,9,nil,nil,3,5}), ps[i], qs[i]).Val.(int))
	}
}
