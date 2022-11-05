package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func reverseList(head *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
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
		rslt := reverseList(datastruct.NewLinkedList[int](v).Head)
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
