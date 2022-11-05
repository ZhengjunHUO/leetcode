package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func deleteMiddle(head *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
	if head.Next == nil {
		return nil
	}

	slow, fast, prev := head, head, &datastruct.LinkedListNode[int]{}

	// when the loop is stop, slow is in the middle of the the list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	// jump over slow
	prev.Next = slow.Next

	return head
}

func main() {
	trees := [][]int{[]int{1,3,4,7,1,2,6}, []int{1,2,3,4}, []int{2,1}}
	for i := range trees {
		rslt := deleteMiddle(datastruct.NewLinkedList[int](trees[i]).Head)
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
