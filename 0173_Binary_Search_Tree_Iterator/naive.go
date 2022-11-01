package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type BSTIterator struct {
	stack	*datastruct.Stack[*godtype.TreeNode]
}

func Constructor(root *godtype.TreeNode) BSTIterator {
	bi := BSTIterator{datastruct.NewStack([]*godtype.TreeNode{})}
	bi.pushNode(root)

	return bi
}

func (this *BSTIterator) Next() int {
	// 中序
	currNode := this.stack.Pop()
	// 后序，注意此处是Right
	this.pushNode(currNode.Right)

	return currNode.Val.(int)
}

func (this *BSTIterator) HasNext() bool {
	return !this.stack.IsEmpty()
}

// 前序
func (this *BSTIterator) pushNode(curr *godtype.TreeNode) {
	for curr != nil {
		this.stack.Push(curr)
		curr = curr.Left
	}
}

func main() {
	obj := Constructor(godtype.NewBTree([]interface{}{7, 3, 15, nil, nil, 9, 20}))
	fmt.Println(obj.Next())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
}
