package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

// 参考@dietpepsi的方案
func largestRectangleArea(heights []int) int {
	// 在末尾添上0，配合后面的算法可在最后让stack元素全部pop出
	list := append(heights, 0)

	mono := datastruct.NewStack([]int{})
	// index指向最后添上的0作为stack的基础，配合后面的算法
	mono.Push(len(list)-1)

	max := 0

	for i := range list {
		// 递增的stack遇到一个小于栈顶的元素，pop掉比它大的元素后插入，使stack始终递增
		// 对每次pop出来的“大”元素进行计算，用变量max保存其中最大值
		for !mono.IsEmpty() && list[i] < list[mono.Peek()] {
			idx := mono.Pop()
			curr := 0
			// 已到初始时插入的垫底元素
			if temp := mono.Peek(); temp == len(list) - 1 {
				curr = list[idx] * i
			}else{
				// 注意i减去的是pop后再peek到的index
				// 因为pop出的高度有效距离最远只能达到其前一个元素的位置为止
				curr = list[idx] * ( i - temp - 1)
			}
			//fmt.Printf("curr idx: %d; poped idx: %d; h&curr: %d %d\n", i, idx, list[idx], curr)
			if curr > max {
				max = curr
			}
		}
		mono.Push(i)
	}

	return max
}

func main() {
	hs := [][]int{[]int{2,1,5,6,2,3}, []int{2,4}}
	for i := range hs {
		fmt.Println(largestRectangleArea(hs[i]))
	}
}
