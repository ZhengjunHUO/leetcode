package main

import (
	"github.com/ZhengjunHUO/godtype"
)

var endNext *godtype.ListNode

// 转置当前链表的前k个元素，类似func见0092 naive2.go
func reverseK(curr *godtype.ListNode, k int) *godtype.ListNode {
	// 如果当前链表不满k个元素返回nil
	if curr == nil {
		return nil
	}

	if k == 1 {
		endNext = curr.Next
		return curr
	}

	end := reverseK(curr.Next, k - 1)
	// 将nil返回到最高层
	if end == nil {
		return nil
	}
	
	curr.Next.Next = curr
	curr.Next = endNext
		
	return end
}

func reverseKGroup(head *godtype.ListNode, k int) *godtype.ListNode {
	curr := head
	// 返回点
	begin := reverseK(head, k)

	for curr.Next != nil {
		temp := endNext
		end := reverseK(endNext, k)
		// 如果返回nil表明后面的元素已经不满k，reverseK没有进行操作
		if end == nil {
			break
		}
		// 反转好的k元素组接到当前元素之后
		curr.Next = end
		// 当前元素前进到新append组的末尾
		curr = temp
	}

	return begin
}

func main() {
	ks := []int{2,3,1}
	for _, k := range ks {
		godtype.PrintList(reverseKGroup(godtype.NewList([]int{1,2,3,4,5}), k))
	}

	godtype.PrintList(reverseKGroup(godtype.NewList([]int{1}), 1))
}
