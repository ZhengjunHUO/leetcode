package main

import (
	"github.com/ZhengjunHUO/goutil/datastruct"
)


func removeNthFromEnd(list *datastruct.LinkedList[int], n int) *datastruct.LinkedList[int] {
	head := list.Head
	if head.Next == nil {
		return &datastruct.LinkedList[int]{}
	}

	curr := head
	mark := &datastruct.LinkedListNode[int]{}
	delay := 0 - n

	for curr.Next != nil {
		delay++
		if delay == 0 {
			mark = head
		}
		if delay > 0 {
			mark = mark.Next
		}

		curr = curr.Next
	}

	if mark.Next != nil {
		mark.Next = mark.Next.Next
	}

	return list
}

func main() {
	l := [][]int{[]int{1,2,3,4,5}, []int{1}, []int{1,2}}
	n := []int{2, 1, 1}

	for i,v := range l {
		list := datastruct.NewLinkedList[int](v)
		removeNthFromEnd(list, n[i]).PrintAll()
	}
}
