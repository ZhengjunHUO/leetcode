package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func lowestCommonAncestor(root *godtype.TreeNode, p, q int) *godtype.TreeNode {
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
		fmt.Println(lowestCommonAncestor(godtype.NewBTree([]interface{}{6,2,8,0,4,7,9,nil,nil,3,5}), ps[i], qs[i]).Val.(int))
	}
}
