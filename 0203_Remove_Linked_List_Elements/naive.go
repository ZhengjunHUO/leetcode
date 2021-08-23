package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func removeElements(head *godtype.ListNode, val int) *godtype.ListNode {
	dummy := &godtype.ListNode{0, head}
	prev, curr := dummy, head

	for curr != nil {
		if curr.Val.(int) == val {
			curr = curr.Next	
			prev.Next = curr
		}else{		
			curr, prev = curr.Next, curr
		}
	}

	return dummy.Next
}

func main() {
	lists := [][]int{[]int{1,2,6,3,4,5,6}, []int{7,7,7,7}, []int{}}
	vals := []int{6,7,1}
	for i := range lists {
		godtype.PrintList(removeElements(godtype.NewList(lists[i]), vals[i]))
	}
}
