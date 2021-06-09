package main

import (
	"fmt"
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
	curr1, curr2 := l1, l2    
	head := &ListNode{}
	curr := head

	val1, val2, carry := 0, 0, 0

	for curr1 != nil || curr2 != nil || carry > 0 {
		if curr1 != nil {
			val1 = curr1.Val
		}else{
			val1 = 0
		}		
		if curr2 != nil {
			val2 = curr2.Val
		}else{
			val2 = 0
		}

		sum := val1 + val2 + carry
		carry = sum/10

		curr.Next = &ListNode{sum%10, nil}
		curr = curr.Next

		if curr1 != nil {
			curr1 = curr1.Next
		}

		if curr2 != nil {
			curr2 = curr2.Next
		}
	}

	return head.Next
}

func main(){
	l1 := newList([]int{2,4,3})
	l2 := newList([]int{5,6,4})
	printList(addTwoNumbers(l1, l2))

	l1 = newList([]int{0})
	l2 = newList([]int{0})
	printList(addTwoNumbers(l1, l2))
	
	l1 = newList([]int{9,9,9,9,9,9,9})
	l2 = newList([]int{9,9,9,9})
	printList(addTwoNumbers(l1, l2))
}
