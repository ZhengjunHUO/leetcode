package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func mergeTwoLists(l1 *godtype.ListNode, l2 *godtype.ListNode) *godtype.ListNode {
	curr := &godtype.ListNode{nil, nil}
	rslt := curr

	for l1 != nil && l2 != nil {
		if l1.Val.(int) < l2.Val.(int) {
			curr.Next = &godtype.ListNode{l1.Val, nil}
			l1 = l1.Next
		}else{
			curr.Next = &godtype.ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 !=nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

   	return rslt.Next
}

func main() {
	godtype.PrintList(mergeTwoLists(godtype.NewList([]int{1,2,4}), godtype.NewList([]int{1,3,4})))
	godtype.PrintList(mergeTwoLists(godtype.NewList([]int{}), godtype.NewList([]int{})))
	godtype.PrintList(mergeTwoLists(godtype.NewList([]int{}), godtype.NewList([]int{0})))
}
