package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func reverse(node1, node2 *godtype.ListNode) *godtype.ListNode {
	var rslt *godtype.ListNode
	if node2.Next != nil {
		rslt = reverse(node2, node2.Next)
	}else{
		rslt = node2
	}

	node1.Next, node2.Next = nil, node1
	return rslt
}

func reverseList(head *godtype.ListNode) *godtype.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rslt := reverse(head, head.Next)
	return rslt
}

func main() {
	lists := [][]int{[]int{1,2,3,4,5}, []int{1,2}, []int{}}
	for _, v := range lists {
		godtype.PrintList(reverseList(godtype.NewList(v)))
	}
}
