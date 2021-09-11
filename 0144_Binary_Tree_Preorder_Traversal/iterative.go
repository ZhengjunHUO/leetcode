package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

// 参考@pavel-shlyk的算法
func preorderTraversal(root *godtype.TreeNode) []int {
	rslt := []int{}
	stack := godtype.NewStack()
	
	for root != nil {
		rslt = append(rslt, root.Val.(int))
		if root.Right != nil {
			stack.Push(root.Right)
		}

		root = root.Left
		if root == nil && !stack.IsEmpty() {
			root = stack.Pop().(*godtype.TreeNode)
		}
	}

	return rslt   
}

func main() {
	ts := [][]interface{}{[]interface{}{1,nil,2,nil,nil,3}, []interface{}{}, []interface{}{1}, []interface{}{1,2}, []interface{}{1,nil,2}}
	for i := range ts {
		fmt.Println(preorderTraversal(godtype.NewBTree(ts[i])))
	}
}
