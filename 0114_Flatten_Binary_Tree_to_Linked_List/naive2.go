package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func flatten(root *godtype.TreeNode)  {
	if root == nil {
		return
	}

	// 前进到当前分支最左下角的三角形，然后处理左下角包含该三角形的上一级三角形以此类推
	flatten(root.Left)
	flatten(root.Right)

	// 处于三角形的顶点，操作左右分支
	originRight, appendPoint := root.Right, root
	// (1) 将左分支变为右分支
	root.Left, root.Right = nil, root.Left

	// (2) 将原右分支接再原左分支的末尾
	for appendPoint.Right != nil {
		appendPoint = appendPoint.Right
	}
	appendPoint.Right = originRight
}

func main() {
	tree := godtype.NewBTree([]interface{}{1,2,5,3,4,nil,6}, 0)
	flatten(tree)
	godtype.PrintBtree(tree)

	tree1 := godtype.NewBTree([]interface{}{1,2,6,3,5,nil,7,4}, 0)
	flatten(tree1)
	godtype.PrintBtree(tree1)
}
