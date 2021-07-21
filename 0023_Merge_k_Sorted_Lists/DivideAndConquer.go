package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func mergeKLists(lists []*godtype.ListNode) *godtype.ListNode {
	size := len(lists)
	if size == 0 {
		return nil
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
func mergeTwoLists(l1 *godtype.ListNode, l2 *godtype.ListNode) *godtype.ListNode {
	curr := &godtype.ListNode{nil, nil}
	rslt := curr

	for l1 != nil && l2 != nil {
		if l1.Val.(int) < l2.Val.(int) {
			curr.Next = &godtype.ListNode{l1.Val, nil}
			l1 = l1.Next
		}else{
			curr.Next = &godtype.ListNode{l2.Val, nil}
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

   	return rslt.Next
}

func main() {
	godtype.PrintList(mergeKLists([]*godtype.ListNode{godtype.NewList([]int{1,4,5}), godtype.NewList([]int{1,3,4}), godtype.NewList([]int{2,6})}))
	godtype.PrintList(mergeKLists([]*godtype.ListNode{}))
	godtype.PrintList(mergeKLists([]*godtype.ListNode{godtype.NewList([]int{})}))
}
