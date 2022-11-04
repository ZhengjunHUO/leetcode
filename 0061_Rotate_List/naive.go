package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func rotateRight(head *datastruct.LinkedListNode[int], k int) *datastruct.LinkedListNode[int] {
	curr, newEnd, size := head, head, 1

	for curr.Next != nil {
		curr = curr.Next
		size += 1
	}

	forward := size - k%size - 1
	for forward > 0 {
		newEnd = newEnd.Next
		forward--
	}

	rslt := newEnd.Next
	newEnd.Next = nil
	curr.Next = head

	return rslt
}

func main() {
	ls := []*datastruct.LinkedList[int]{datastruct.NewLinkedList[int]([]int{1,2,3,4,5}), datastruct.NewLinkedList[int]([]int{0,1,2})}
	ks := []int{2,4}
	for i := range ls {
		rslt := rotateRight(ls[i].Head, ks[i])
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
