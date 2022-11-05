package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func reverse(head *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
	curr := head
	var prev, next *datastruct.LinkedListNode[int]

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}

	return prev
}

func isPalindrome(head *datastruct.LinkedListNode[int]) bool {
	// 寻找链表中点，参照0876
	mid, fast := head, head
	// 对应长度为偶数和奇数两种情况
	for fast != nil && fast.Next != nil {
		mid, fast = mid.Next, fast.Next.Next
	}

	// 长度为奇数，此时mid在正中心，需要再前进一格移到中点偏右的位置
	if fast != nil {
		mid = mid.Next
	}

	// 反转mid以后的链表，r为原链表尾现在为反转后的首部
	l, r := head, reverse(mid)
	// r的尾部为nil，且右边反转的链表长度等于左边长度或左边长度-1
	for r != nil {
		// 两头向中间逐个比较
		if l.Val != r.Val {
			return false
		}
		l, r = l.Next, r.Next
	}

	return true
}

func main() {
	ls := [][]int{[]int{1,2,2,1},[]int{1,2},[]int{1,2,3,4,5},[]int{1,2,3,2,1}}
	for i := range ls {
		fmt.Println(isPalindrome(datastruct.NewLinkedList(ls[i]).Head))
	}
}
