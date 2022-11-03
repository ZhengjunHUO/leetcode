package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func swapPairs(head *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	rslt, curr := head.Next, head
	var prev *datastruct.LinkedListNode[int]

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
	ls := []*datastruct.LinkedList[int]{
		datastruct.NewLinkedList[int]([]int{1,2,3,4}),
		datastruct.NewLinkedList[int]([]int{1,2,3,4,5,6,7}),
		datastruct.NewLinkedList[int]([]int{}),
		datastruct.NewLinkedList[int]([]int{1}),
	}

	for i := range ls {
		rslt := swapPairs(ls[i].Head)
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
