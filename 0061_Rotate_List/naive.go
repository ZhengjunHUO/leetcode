package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func rotateRight(head *godtype.ListNode, k int) *godtype.ListNode {
	curr, newEnd, size := head, head, 1

	for curr.Next != nil {
		curr = curr.Next
		size += 1
	}

	forward := size - k%size - 1	
	for forward > 0 {
		newEnd = newEnd.Next
		forward--
	}

	rslt := newEnd.Next
	newEnd.Next = nil
	curr.Next = head
    
	return rslt
}

func main() {
	godtype.PrintList(rotateRight(godtype.NewList([]int{1,2,3,4,5}), 2))
	godtype.PrintList(rotateRight(godtype.NewList([]int{0,1,2}), 4))
}
