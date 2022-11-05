package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func reverse(node1, node2 *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
	var rslt *datastruct.LinkedListNode[int]
	if node2.Next != nil {
		rslt = reverse(node2, node2.Next)
	}else{
		rslt = node2
	}

	node1.Next, node2.Next = nil, node1
	return rslt
}

func reverseList(head *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
	if head == nil || head.Next == nil {
		return head
	}

	rslt := reverse(head, head.Next)
	return rslt
}

func main() {
	lists := [][]int{[]int{1,2,3,4,5}, []int{1,2}, []int{}}
	for _, v := range lists {
		rslt := reverseList(datastruct.NewLinkedList[int](v).Head)
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
