package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func reverseList(head *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
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
		rslt := reverseList(datastruct.NewLinkedList[int](v).Head)
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
