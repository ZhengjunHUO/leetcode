package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func postorderTraversal(root *godtype.TreeNode) []int {
	rslt := []int{}
	stack := datastruct.NewStack([]*godtype.TreeNode{})

	if root == nil {
		return rslt
	}
	stack.Push(root)

	for !stack.IsEmpty() {
		// 当前节点的值放在列表最后
		curr := stack.Pop()
		rslt = append([]int{curr.Val.(int)}, rslt...)

		// 按这个顺序下一轮会先Pop右结点，按先右后左处理
		if curr.Left != nil {
			stack.Push(curr.Left)
		}
		if curr.Right != nil {
			stack.Push(curr.Right)
		}
	}

	return rslt
}

func main() {
	ts := [][]interface{}{[]interface{}{1,nil,2,nil,nil,3}, []interface{}{}, []interface{}{1}, []interface{}{1,2}, []interface{}{1,nil,2}}
	for i := range ts {
		fmt.Println(postorderTraversal(godtype.NewBTree(ts[i])))
	}
}
