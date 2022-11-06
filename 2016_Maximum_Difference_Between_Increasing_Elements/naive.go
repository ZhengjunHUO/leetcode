package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maximumDifference(nums []int) int {
	mono := datastruct.NewStack[int]([]int{})
	ret := -1

	for i := range nums {
		for !mono.IsEmpty() && nums[i] <= mono.Peek() {
			mono.Pop()
		}

		mono.Push(nums[i])
		if mono.Size() > 1 {
			ret = max(ret, nums[i] - mono.Elems[0])
		}
	}

	return ret
}

func main() {
	nums := [][]int{[]int{7,1,5,4}, []int{9,4,3,2}, []int{1,5,2,10}}
	for i := range nums {
		fmt.Println(maximumDifference(nums[i]))
	}
}
