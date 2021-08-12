package main

import (
	"github.com/ZhengjunHUO/godtype"
)

/*
  当前结点的值为原来值加上右子树的return值
  再把update后的结点值传给左子树
  最后向上层返回左子树递归的结果
*/
func parse(root *godtype.TreeNode, x int) int {
	if root == nil {
		return x
	}

	root.Val = root.Val.(int) + parse(root.Right, x)
	return parse(root.Left, root.Val.(int))
}

func convertBST(root *godtype.TreeNode) *godtype.TreeNode {
	// 最右下角的值不需要修改，所以初始传递值为0
	parse(root, 0)
	return root
}

func main() {
	trees := [][]interface{}{[]interface{}{4,1,6,0,2,5,7,nil,nil,nil,3,nil,nil,nil,8}, []interface{}{0,nil,1}, []interface{}{1,0,2}, []interface{}{3,2,4,1}}
	for i := range trees {
		godtype.PrintBTreeBFS(convertBST(godtype.NewBTree(trees[i])))
	}	
}
