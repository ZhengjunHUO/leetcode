package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func mergeNodes(head *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
	ret := &datastruct.LinkedListNode[int]{}
	curr := ret

	var curr_sum int
	head = head.Next
	for head != nil {
		if head.Val == 0 {
			curr.Next = &datastruct.LinkedListNode[int]{curr_sum, nil}
			curr = curr.Next
			curr_sum = 0
		}else{
			curr_sum += head.Val
		}

		head = head.Next
	}

	return ret.Next
}

func main() {
	list := [][]int{[]int{0,3,1,0,4,5,2,0}, []int{0,1,0,3,0,2,2,0}}
	for i := range list {
		rslt := mergeNodes(datastruct.NewLinkedList(list[i]).Head)
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
