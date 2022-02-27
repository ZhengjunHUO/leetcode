package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func pairSum(head *godtype.ListNode) int {
	// 如果只有两个节点，直接返回和
	if head.Next.Next == nil {
		return head.Val.(int) + head.Next.Val.(int)
	}

	// fast每回合往前走两步，slow每回合移动一格
	// prev, next记录slow的前一个和后一个节点
	slow, next, fast := head, head.Next, head
	var prev *godtype.ListNode
	var ret int

	// 快指针走完链表时，慢指针会正好停在中间偏右的位置(偶数个节点)
	for fast != nil {
		fast = fast.Next.Next

		prev = slow
		slow = next
		next = slow.Next
		// 翻转指针方向
		slow.Next = prev
	}
	// 前半个链表已被翻转，将head的Next置nil消除环
	head.Next = nil

	// left负责遍历翻转后的左半边
	left := slow.Next
	// 修复slow的指针重新指向右边
	slow.Next = next
	// 同时遍历两边
	for left != nil && slow != nil {
		//fmt.Printf("%d + %d\n", left.Val.(int), slow.Val.(int))
		ret = max(ret, left.Val.(int)+slow.Val.(int))
		left, slow = left.Next, slow.Next
	}

	return ret
}

func main() {
	heads := [][]int{[]int{5,4,2,1}, []int{4,2,2,3}, []int{1,100000}}
	for i := range heads {
		fmt.Println(pairSum(godtype.NewList(heads[i])))
	}
}
