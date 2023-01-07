package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

// Complete二叉树至少有一边是Perfect二叉树(2^h-1个结点)
// 时间复杂度为O(logN*logN)
func countNodes(root *graph.TreeNode) int {
	if root == nil {
		return 0
	}

	l, r, ldep, rdep := root, root, 0, 0
	for l.Left != nil {
		l = l.Left
		ldep++
	}

	for r.Right != nil {
		r = r.Right
		rdep++
	}

	// 找到了Perfect子二叉树，返回2^h - 1
	// ldep从0开始计，所以实际树高为ldep+1，考虑到2开始<<运算已经算是2^1
	// 所以+1-1互相抵消，使用2<<ldep - 1来计算
	if ldep == rdep {
		return 2<<ldep - 1
	}

	// 这一支不是Perfect二叉树，按传统方法计算
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func main() {
	tree := [][]interface{}{[]interface{}{1,2,3,4,5,6}, []interface{}{}, []interface{}{1}}
	for i := range tree {
		fmt.Println(countNodes(graph.NewBTree(tree[i])))
	}
}
