package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

// fast每次向前走两格，slow走一格，如果有环的话两者会相遇
func hasCycle(head *godtype.ListNode) bool {
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
	list := godtype.NewList([]int{3,2,0,-4})
	target := list
	for target.Next != nil {
		target = target.Next	
	}
	target.Next = list.Next
	fmt.Println(hasCycle(list))

	list1 := godtype.NewList([]int{1,2})
	list1.Next.Next = list1
	fmt.Println(hasCycle(list1))

	fmt.Println(hasCycle(godtype.NewList([]int{1})))
}
