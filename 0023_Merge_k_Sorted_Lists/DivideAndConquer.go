package main

import (
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func mergeKLists(lists []*datastruct.LinkedList[int]) *datastruct.LinkedList[int] {
	size := len(lists)
	if size == 0 {
		return &datastruct.LinkedList[int]{}
	}

	itv := 1
	for itv < size {
		for i:=0; i<size-itv; i+=2*itv {
			lists[i] = mergeTwoLists(lists[i], lists[i+itv])
		}
		itv *= 2
	}

	return lists[0]
}

// 拷贝自0021_Merge_Two_Sorted_Lists/naive.go
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
	mergeKLists([]*datastruct.LinkedList[int]{datastruct.NewLinkedList([]int{1,4,5}), datastruct.NewLinkedList([]int{1,3,4}), datastruct.NewLinkedList([]int{2,6})}).PrintAll()
	mergeKLists([]*datastruct.LinkedList[int]{}).PrintAll()
	mergeKLists([]*datastruct.LinkedList[int]{datastruct.NewLinkedList([]int{})}).PrintAll()
}
