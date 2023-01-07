package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func isSameTree(p *graph.TreeNode, q *graph.TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	ps := [][]interface{}{[]interface{}{1,2,3}, []interface{}{1,2}, []interface{}{1,2,1}}
	qs := [][]interface{}{[]interface{}{1,2,3}, []interface{}{1,nil,2}, []interface{}{1,1,2}}
	for i := range ps {
		fmt.Println(isSameTree(graph.NewBTree(ps[i]), graph.NewBTree(qs[i])))
	}
}
