package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

var endNext *datastruct.LinkedListNode[int]

// 转置当前链表的前rest个元素
func reverse(curr *datastruct.LinkedListNode[int], rest int) *datastruct.LinkedListNode[int] {
	// 区间尾部，这个结点将会被层层return到调用入口
	if rest == 1 {
		endNext = curr.Next
		return curr
	}

	// 前进到区间尾部
	end := reverse(curr.Next, rest - 1)

	// 转向
	curr.Next.Next = curr
	// 不用指向nil，因为中间的结点的Next会在上一层调用重定向（会被覆盖）
	// 而最后一层结点的Next正好需要指向endNext
	curr.Next = endNext

	return end
}

func reverseBetween(head *datastruct.LinkedListNode[int], left int, right int) *datastruct.LinkedListNode[int] {
	// 区间开始处，进入reverse()，功能为转置链表前right个元素
	if left == 1 {
		return reverse(head, right)
	}

	// 前进到区间开始处
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
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
