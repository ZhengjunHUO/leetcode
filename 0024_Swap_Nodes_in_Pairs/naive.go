package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func swapPairs(head *godtype.ListNode) *godtype.ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	rslt, curr := head.Next, head
	var prev *godtype.ListNode

	for curr != nil && curr.Next != nil {
		temp := curr.Next
		if prev != nil {
			prev.Next = temp
		}
		curr.Next, temp.Next = temp.Next, curr
		curr, prev = curr.Next, curr	
	}

	return rslt
}

func main() {
	godtype.PrintList(swapPairs(godtype.NewList([]int{1,2,3,4})))
	godtype.PrintList(swapPairs(godtype.NewList([]int{1,2,3,4,5,6,7})))
	godtype.PrintList(swapPairs(godtype.NewList([]int{})))
	godtype.PrintList(swapPairs(godtype.NewList([]int{1})))
}
