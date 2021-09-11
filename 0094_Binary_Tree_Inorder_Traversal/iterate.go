package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

// 借鉴0173_Binary_Search_Tree_Iterator的思路
func inorderTraversal(root *godtype.TreeNode) []int {
	rslt := []int{}
	stack := godtype.NewStack()

	// 添加当前结点的所有左子结点
	push := func(curr *godtype.TreeNode) {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
	}

	push(root)
		
	for !stack.IsEmpty() {
		root = stack.Pop().(*godtype.TreeNode)	
		// 中序输出
		rslt = append(rslt, root.Val.(int))
		// 加入右结点及其所有左子结点
		push(root.Right)
	}

	return rslt   
}

func main() {
	ts := [][]interface{}{[]interface{}{1,nil,2,nil,nil,3}, []interface{}{}, []interface{}{1}, []interface{}{1,2}, []interface{}{1,nil,2}}
	for i := range ts {
		fmt.Println(inorderTraversal(godtype.NewBTree(ts[i])))
	}
}
