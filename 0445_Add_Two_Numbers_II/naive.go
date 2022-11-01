package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func newList(l []int) *ListNode {
	head := &ListNode{l[0], nil}
	n := len(l)

	if n == 1 {
		return head
	}

	curr := head
	for i:=1; i<n; i++ {
		curr.Next = &ListNode{l[i], nil}
		curr = curr.Next
	}
	return head
}

func printList(node *ListNode) {
	rslt := []int{}
	curr := node
	for {
		rslt = append(rslt, curr.Val)

		if curr.Next == nil {
			break
		}

		curr = curr.Next
	}

	fmt.Println(rslt)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c1, c2 := l1, l2
	s1, s2 := datastruct.NewStack([]int{}), datastruct.NewStack([]int{})

	var curr *ListNode
	v1, v2, carry := 0, 0, 0

	for c1 != nil {
		s1.Push(c1.Val)
		c1 = c1.Next
	}
	for c2 != nil {
		s2.Push(c2.Val)
		c2 = c2.Next
	}

	for !s1.IsEmpty() || !s2.IsEmpty() || carry > 0 {
		if !s1.IsEmpty() {
			v1 = s1.Pop()
		}else{
			v1 = 0
		}
		if !s2.IsEmpty() {
			v2 = s2.Pop()
		}else{
			v2 = 0
		}

		sum := v1 + v2 + carry
		carry = sum/10

		// 反向制造链表
		prev := &ListNode{sum%10, curr}
		curr = prev
	}

	return curr
}

func main(){
	l1 := newList([]int{7,2,4,3})
	l2 := newList([]int{5,6,4})
	printList(addTwoNumbers(l1, l2))

	l1 = newList([]int{2,4,3})
	l2 = newList([]int{5,6,4})
	printList(addTwoNumbers(l1, l2))

	l1 = newList([]int{0})
	l2 = newList([]int{0})
	printList(addTwoNumbers(l1, l2))
}
