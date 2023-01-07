package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func lowestCommonAncestor(root *graph.TreeNode, p, q int) *graph.TreeNode {
	if root == nil {
		return nil
	}

	if p > q {
		p, q = q, p
	}

	if root.Val.(int) >= p && root.Val.(int) <= q {
		//fmt.Printf("Found %v\n", root.Val.(int))
		return root
	}

	var rslt *graph.TreeNode
	if root.Val.(int) < p {
		//fmt.Printf("%v: search right %v\n", root.Val.(int), root.Right.Val.(int))
		rslt = lowestCommonAncestor(root.Right, p, q)
	}else{
		//fmt.Printf("%v: search left %v\n", root.Val.(int), root.Left.Val.(int))
		rslt = lowestCommonAncestor(root.Left, p, q)
	}

	return rslt
}

func main() {
	ps := []int{2,2,0}
	qs := []int{8,4,5}

	for i := range ps {
		fmt.Println(lowestCommonAncestor(graph.NewBTree([]interface{}{6,2,8,0,4,7,9,nil,nil,3,5}), ps[i], qs[i]).Val.(int))
	}
}
