package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func isSubtree(root *graph.TreeNode, subRoot *graph.TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	// 如果结点值不同，检查两个分支是否能匹配子树
	if root.Val.(int) != subRoot.Val.(int) {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}

	// 结点值相同的情况下，可以假定两颗树的根已经匹配，需比较左右分支是否都相同
	// 或者子树仍可能存在于左或右分支中
	return (isSubtree(root.Left, subRoot.Left) && isSubtree(root.Right, subRoot.Right)) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func main() {
	roots := [][]interface{}{[]interface{}{3,4,5,1,2}, []interface{}{3,4,5,1,2,nil,nil,nil,nil,0}}
	subroots := [][]interface{}{[]interface{}{4,1,2}, []interface{}{4,1,2}}

	for i := range roots {
		fmt.Println(isSubtree(graph.NewBTree(roots[i]), graph.NewBTree(subroots[i])))
	}
}
