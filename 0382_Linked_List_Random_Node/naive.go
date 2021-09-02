package main

import (
	"fmt"
	"math/rand"
	"github.com/ZhengjunHUO/godtype"
)

type Solution struct {
	list	*godtype.ListNode    
}

func Constructor(head *godtype.ListNode) Solution {
	return Solution{
		list: head,
	}
}

func (this *Solution) GetRandom() int {
	curr := this.list
	rslt, i := 0, 0

	// 遍历链表
	for curr != nil {
		i++
		// 每一轮以1/i的概率来更新结果
		// 最终结果能够达成uniform random
		if rand.Intn(i)	== 0 {
			rslt = curr.Val.(int)
		}
		curr = curr.Next
	}

	return rslt
}


func main() {
	obj := Constructor(godtype.NewList([]int{1,2,3,4,5}))
	dict := make(map[int]int)
	for i:=0; i<10000; i++ {
		dict[obj.GetRandom()]++
	}

	fmt.Println(dict)
}
