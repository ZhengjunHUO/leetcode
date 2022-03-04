package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func deleteMiddle(head *godtype.ListNode) *godtype.ListNode {
	if head.Next == nil {
		return nil
	}

	slow, fast, prev := head, head, &godtype.ListNode{}

	// when the loop is stop, slow is in the middle of the the list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	// jump over slow
	prev.Next = slow.Next

	return head
}

func main() {
	trees := [][]int{[]int{1,3,4,7,1,2,6}, []int{1,2,3,4}, []int{2,1}}
	for i := range trees {
		godtype.PrintList(deleteMiddle(godtype.NewList(trees[i])))
	}
}
