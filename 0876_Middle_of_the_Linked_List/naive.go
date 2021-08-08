package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func middleNode(head *godtype.ListNode) *godtype.ListNode {
	fast, slow := head, head
	
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func main() {
	list := [][]int{[]int{1,2,3,4,5}, []int{1,2,3,4,5,6}}
	for i := range list {
		fmt.Println(middleNode(godtype.NewList(list[i])).Val)
	}	

}
