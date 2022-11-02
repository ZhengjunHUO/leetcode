package main

import (
	"fmt"
	"math/rand"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type Solution struct {
	list	*datastruct.LinkedList[int]
}

func Constructor(head *datastruct.LinkedList[int]) Solution {
	return Solution{
		list: head,
	}
}

func (this *Solution) GetRandom() int {
	curr := this.list.Head
	rslt, i := 0, 0

	// 遍历链表
	for curr != nil {
		i++
		// 每一轮以1/i的概率来更新结果
		// 最终结果能够达成uniform random
		if rand.Intn(i)	== 0 {
			rslt = curr.Val
		}
		curr = curr.Next
	}

	return rslt
}


func main() {
	obj := Constructor(datastruct.NewLinkedList([]int{1,2,3,4,5}))
	dict := make(map[int]int)
	for i:=0; i<10000; i++ {
		dict[obj.GetRandom()]++
	}

	fmt.Println(dict)
}
