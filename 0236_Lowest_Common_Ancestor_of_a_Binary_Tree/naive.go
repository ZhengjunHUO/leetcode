package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

var lca *godtype.TreeNode

func lowestCommonAncestor(root *godtype.TreeNode, p, q int) *godtype.TreeNode {
	if root == nil {
		return nil
	}	

	found := 0	

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	
	if lca != nil {
		return lca		
	}

	if left != nil {
		found += 1
	}
	if right != nil {
		found += 1
	}
	if root.Val.(int) == p || root.Val.(int) == q {
		found += 1	
	}

	if found == 0 {
		return nil
	}
	if found > 1 {
		lca = root
		return lca
	} 

	return root
}

func main() {
	ps := []int{5,5}
	qs := []int{1,4}

	for i := range ps {
		lca = nil
		fmt.Println(lowestCommonAncestor(godtype.NewBTree([]interface{}{3,5,1,6,2,0,8,nil,nil,7,4}), ps[i], qs[i]).Val.(int))
	}
}
