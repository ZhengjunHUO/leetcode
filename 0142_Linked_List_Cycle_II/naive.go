package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func detectCycle(head *datastruct.LinkedListNode[int]) *datastruct.LinkedListNode[int] {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		// 在0141基础上改进，在快慢指针汇合后，令其中一个返回头部
		// 然后以同样速度前进，再次相遇时即为寻找的点
		if fast == slow {
			slow = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			return slow
		}
	}

	return nil
}

func main() {
	list := datastruct.NewLinkedList[int]([]int{3,2,0,-4})
	target := list.Head
	for target.Next != nil {
		target = target.Next
	}
	target.Next = list.Head.Next
	fmt.Println(detectCycle(list.Head).Val)

	list1 := datastruct.NewLinkedList[int]([]int{1,2})
	list1.Head.Next.Next = list1.Head
	fmt.Println(detectCycle(list1.Head).Val)

	fmt.Println(detectCycle(datastruct.NewLinkedList[int]([]int{1}).Head))
}
