package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

// fast每次向前走两格，slow走一格，如果有环的话两者会相遇
func hasCycle(head *datastruct.LinkedListNode[int]) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

func main() {
	list := datastruct.NewLinkedList[int]([]int{3,2,0,-4})
	target := list.Head
	for target.Next != nil {
		target = target.Next
	}
	target.Next = list.Head.Next
	fmt.Println(hasCycle(list.Head))

	list1 := datastruct.NewLinkedList[int]([]int{1,2})
	list1.Head.Next.Next = list1.Head
	fmt.Println(hasCycle(list1.Head))

	fmt.Println(hasCycle(datastruct.NewLinkedList[int]([]int{1}).Head))
}
