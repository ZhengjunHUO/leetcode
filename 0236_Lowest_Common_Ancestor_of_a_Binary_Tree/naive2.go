package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func lowestCommonAncestor(root *godtype.TreeNode, p, q int) *godtype.TreeNode {
	if root == nil || root.Val.(int) == p || root.Val.(int) == q {
		return root
	}	

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	// 此时left和right都不是nil，找到lca
	return root
}

func main() {
	ps := []int{5,5}
	qs := []int{1,4}

	for i := range ps {
		fmt.Println(lowestCommonAncestor(godtype.NewBTree([]interface{}{3,5,1,6,2,0,8,nil,nil,7,4}), ps[i], qs[i]).Val.(int))
	}
}
