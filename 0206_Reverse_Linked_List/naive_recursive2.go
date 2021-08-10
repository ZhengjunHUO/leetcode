package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func reverseList(head *godtype.ListNode) *godtype.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rslt := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return rslt
}

func main() {
	lists := [][]int{[]int{1,2,3,4,5}, []int{1,2}, []int{}}
	for _, v := range lists {
		godtype.PrintList(reverseList(godtype.NewList(v)))
	}
}
