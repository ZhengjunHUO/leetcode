package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func detectCycle(head *godtype.ListNode) *godtype.ListNode {
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
	list := godtype.NewList([]int{3,2,0,-4})
	target := list
	for target.Next != nil {
		target = target.Next	
	}
	target.Next = list.Next
	fmt.Println(detectCycle(list).Val)

	list1 := godtype.NewList([]int{1,2})
	list1.Next.Next = list1
	fmt.Println(detectCycle(list1).Val)

	fmt.Println(detectCycle(godtype.NewList([]int{1})))
}
