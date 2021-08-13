package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func findRightMin (root *godtype.TreeNode) *godtype.TreeNode {
	for root.Left != nil {
		root = root.Left
	}

	return root
}

func deleteNode(root *godtype.TreeNode, key int) *godtype.TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val.(int) {
		root.Left = deleteNode(root.Left, key)	
	}else if key > root.Val.(int) {
		root.Right = deleteNode(root.Right, key)
	}else{
		// 如果目标是叶子结点，直接移除
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// 如果目标有一个子结点，则跳过自己，让自己的父结点直接指向自己的子结点	
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 如果目标有两个子结点，则取右子树中最小的结点的值，并让右子树删除该结点
		root.Val = findRightMin(root.Right).Val
		root.Right = deleteNode(root.Right, root.Val.(int))
	}

	return root
}

func main() {
	lists := [][]interface{}{[]interface{}{5,3,6,2,4,nil,7}, []interface{}{5,3,6,2,4,nil,7}, []interface{}{}, []interface{}{5,2,6,1,4,nil,7,nil,nil,3}}
	keys := []int{3, 0, 0, 2}
	for i := range lists {
		tree := godtype.NewBTree(lists[i])
		fmt.Printf("Before: ")
		godtype.PrintBTreeBFS(tree)
		fmt.Printf("After:  ")
		godtype.PrintBTreeBFS(deleteNode(tree, keys[i]))
	}
}
