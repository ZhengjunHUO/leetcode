package main

import (
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func mergeTwoLists(list1 *datastruct.LinkedList[int], list2 *datastruct.LinkedList[int]) *datastruct.LinkedList[int] {
	curr := &datastruct.LinkedListNode[int]{0, nil}
	rslt := curr

	l1, l2 := list1.Head, list2.Head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = &datastruct.LinkedListNode[int]{l1.Val, nil}
			l1 = l1.Next
		}else{
			curr.Next = &datastruct.LinkedListNode[int]{l2.Val, nil}
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

	return &datastruct.LinkedList[int]{Head: rslt.Next, Tail: curr.Next}
}

func main() {
	mergeTwoLists(datastruct.NewLinkedList([]int{1,2,4}), datastruct.NewLinkedList([]int{1,3,4})).PrintAll()
	mergeTwoLists(datastruct.NewLinkedList([]int{}), datastruct.NewLinkedList([]int{})).PrintAll()
	mergeTwoLists(datastruct.NewLinkedList([]int{}), datastruct.NewLinkedList([]int{0})).PrintAll()
}
