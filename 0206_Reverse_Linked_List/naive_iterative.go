package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func reverseList(head *godtype.ListNode) *godtype.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr, next := head, head.Next
	curr.Next = nil
	
	for next != nil {
		nextnext := next.Next
		next.Next = curr
		curr, next = next, nextnext
	}

	return curr	
}

func main() {
	lists := [][]int{[]int{1,2,3,4,5}, []int{1,2}, []int{}}
	for _, v := range lists {
		godtype.PrintList(reverseList(godtype.NewList(v)))
	}
}
