package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

var endNext *datastruct.LinkedListNode[int]

// 转置当前链表的前k个元素，类似func见0092 naive2.go
func reverseK(curr *datastruct.LinkedListNode[int], k int) *datastruct.LinkedListNode[int] {
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

func reverseKGroup(head *datastruct.LinkedListNode[int], k int) *datastruct.LinkedListNode[int] {
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
		list := datastruct.NewLinkedList[int]([]int{1,2,3,4,5})
		rslt := reverseKGroup(list.Head, k)
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}

	rslt := reverseKGroup(datastruct.NewLinkedList([]int{1}).Head, 1)
	for rslt != nil {
		fmt.Printf("%v, ", rslt.Val)
		rslt = rslt.Next
	}
	fmt.Println()
}
