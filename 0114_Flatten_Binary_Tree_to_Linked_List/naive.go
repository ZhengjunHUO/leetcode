package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
)

// 树从右向左拉直整合
func flattenbis(curr, appendR *graph.TreeNode) {
	// 叶子结点负责append右边已经拉直的邻居结点
	if curr.Left == nil && curr.Right == nil {
		if appendR != nil {
			curr.Right = appendR
		}
		return
	}

	if curr.Right != nil {
		flattenbis(curr.Right, appendR)
	}

	if curr.Left != nil {
		// 如果只有左分支，则appendR由其传递下去
		if curr.Right == nil {
			flattenbis(curr.Left, appendR)
		// appendR已由右分支整合，左分支整合当前右分支
		}else{
			flattenbis(curr.Left, curr.Right)
		}
		// 整合结束后根据题目要求将左分支变成右分支
		curr.Right = curr.Left
		curr.Left = nil
	}
}

func flatten(root *graph.TreeNode)  {
	flattenbis(root, nil)
}

func main() {
	tree := graph.NewBTree([]interface{}{1,2,5,3,4,nil,6})
	flatten(tree)
	graph.PrintBTreeDFS(os.Stdout, tree)

	tree1 := graph.NewBTree([]interface{}{1,2,6,3,5,nil,7,4})
	flatten(tree1)
	graph.PrintBTreeDFS(os.Stdout, tree1)
}
