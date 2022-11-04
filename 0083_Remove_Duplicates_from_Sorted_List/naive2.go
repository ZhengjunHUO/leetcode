package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

// 修改指针，跳过重复结点，有内存泄漏?
func deleteDuplicates(head *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
	p1, p2 := head, head

	for p2 != nil {
		if p1.Val != p2.Val {
			p1.Next = p2
			p1 = p1.Next
		}

		p2 = p2.Next
	}

	prev, next := p1, p1.Next
	for next != nil {
		prev.Next = nil
		prev, next = next, next.Next
	}

	return head
}

func main() {
	lists := [][]int{[]int{1,1,2}, []int{1,1,2,3,3}}
	for i := range lists {
		list := datastruct.NewLinkedList[int](lists[i])
		rslt := deleteDuplicates(list.Head)
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
