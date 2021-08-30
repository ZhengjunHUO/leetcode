package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

type BSTIterator struct {
	stack	*godtype.Stack
}

func Constructor(root *godtype.TreeNode) BSTIterator {
	bi := BSTIterator{godtype.NewStack()}
	bi.pushNode(root)
	
	return bi	
}

func (this *BSTIterator) Next() int {
	// 中序
	currNode := this.stack.Pop().(*godtype.TreeNode)
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
