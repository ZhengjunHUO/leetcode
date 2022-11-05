package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func removeElements(head *datastruct.LinkedListNode[int], val int) *datastruct.LinkedListNode[int] {
	dummy := &datastruct.LinkedListNode[int]{0, head}
	prev, curr := dummy, head

	for curr != nil {
		if curr.Val == val {
			curr = curr.Next
			prev.Next = curr
		}else{
			curr, prev = curr.Next, curr
		}
	}

	return dummy.Next
}

func main() {
	lists := [][]int{[]int{1,2,6,3,4,5,6}, []int{7,7,7,7}, []int{}}
	vals := []int{6,7,1}
	for i := range lists {
		rslt := removeElements(datastruct.NewLinkedList[int](lists[i]).Head, vals[i])
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
