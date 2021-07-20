package main

import (
	"github.com/ZhengjunHUO/godtype"
)


func removeNthFromEnd(head *godtype.ListNode, n int) *godtype.ListNode {
	if head.Next == nil {
		return nil
	}

	curr := head
	mark := &godtype.ListNode{}
	delay := 0 - n

	for curr.Next != nil {
		delay++	
		if delay == 0 {
			mark = head
		}
		if delay > 0 {
			mark = mark.Next
		}

		curr = curr.Next
	}

	if mark.Next != nil {
		mark.Next = mark.Next.Next
	}
	
	return head
}

func main() {
	l := [][]int{[]int{1,2,3,4,5}, []int{1}, []int{1,2}}
	n := []int{2, 1, 1}

	for i,v := range l {
		list := godtype.NewList(v)
		godtype.PrintList(removeNthFromEnd(list, n[i]))
	}
}
