package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func mergeNodes(head *godtype.ListNode) *godtype.ListNode {
	ret := &godtype.ListNode{}
	curr := ret

	var curr_sum int
	head = head.Next
	for head != nil {
		if head.Val.(int) == 0 {
			curr.Next = &godtype.ListNode{curr_sum, nil}
			curr = curr.Next
			curr_sum = 0
		}else{
			curr_sum += head.Val.(int)
		}

		head = head.Next
	}

	return ret.Next
}

func main() {
	list := [][]int{[]int{0,3,1,0,4,5,2,0}, []int{0,1,0,3,0,2,2,0}}
	for i := range list {
		godtype.PrintList(mergeNodes(godtype.NewList(list[i])))
	}
}
