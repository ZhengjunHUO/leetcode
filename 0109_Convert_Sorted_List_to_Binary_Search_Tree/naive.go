package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/datastruct"
	"github.com/ZhengjunHUO/goutil/graph"
)

var curr *datastruct.LinkedListNode[int]

func build(l ,r int) *graph.TreeNode {
	if l > r {
		return nil
	}

	mid := l + (r-l)/2

	lc := build(l, mid-1)

	// Inorder traversal
	this := &graph.TreeNode{
		Val:	curr.Val,
		Left:	lc,
		Right:	nil,
	}
	curr = curr.Next

	this.Right = build(mid+1, r)

	return this
}

func sortedListToBST(head *datastruct.LinkedListNode[int]) *graph.TreeNode {
	if head == nil {
		return nil
	}

	size, temp := 0, head
	for temp != nil {
		size++
		temp = temp.Next
	}

	curr = head

	return build(0, size-1)
}

func main() {
	heads := [][]int{[]int{-10,-3,0,5,9}, []int{}, []int{0}, []int{1,3}}
	for i := range heads {
		graph.PrintBTreeBFS(os.Stdout, sortedListToBST(datastruct.NewLinkedList[int](heads[i]).Head))
	}
}
