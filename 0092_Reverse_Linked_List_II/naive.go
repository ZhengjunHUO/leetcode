package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func reverse(curr *datastruct.LinkedListNode[int], rest int) (*datastruct.LinkedListNode[int], *datastruct.LinkedListNode[int]) {
	if rest == 0 {
		return curr, curr.Next
	}

	end, endNext := reverse(curr.Next, rest - 1)

	curr.Next.Next = curr
	curr.Next = nil

	return end, endNext
}

func reverseBetween(head *datastruct.LinkedListNode[int], left int, right int) *datastruct.LinkedListNode[int] {
	// 转置区间之前的一个元素，其Next将指向转置后的区间头
	var prebegin *datastruct.LinkedListNode[int]
	// 转置区间的首部
	begin := head

	left, rest := left - 1, right - left

	for left > 0 {
		prebegin = begin
		begin = begin.Next
		left--
	}

	// 转置前区间尾部（新首部），及其原来指向的下一个元素（可能是nil）
	end, endNext := reverse(begin, rest)
	// 接在新尾部上
	begin.Next = endNext

	if prebegin != nil {
		prebegin.Next = end
		return head
	}

	// 区间从原链表头开始，无须接上prebegin
	return end
}

func main() {
	lists := [][]int{[]int{1,2,3,4,5}, []int{5}}
	lefts := []int{2, 1}
	rights := []int{4, 1}

	for i := range lists {
		list := datastruct.NewLinkedList[int](lists[i])
		rslt := reverseBetween(list.Head, lefts[i], rights[i])
		for rslt != nil {
			fmt.Printf("%v, ", rslt.Val)
			rslt = rslt.Next
		}
		fmt.Println()
	}
}
