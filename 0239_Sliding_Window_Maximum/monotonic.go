package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums) 

	mq := godtype.NewMonotonicQueue()
	rslt := make([]int, n-k+1)

	for i := range nums{
		if i < k - 1 {
			mq.Push(nums[i])
			continue
		}

		mq.Push(nums[i])
		rslt[i-k+1] = mq.Max().(int)
		mq.Pop(nums[i-k+1])
	}

	return rslt	
}

func main() {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))

	nums = []int{1}
	k = 1
	fmt.Println(maxSlidingWindow(nums, k))

	nums = []int{1, -1}
	k = 1
	fmt.Println(maxSlidingWindow(nums, k))

	nums = []int{9, 11}
	k = 2
	fmt.Println(maxSlidingWindow(nums, k))

	nums = []int{4, -2}
	k = 2
	fmt.Println(maxSlidingWindow(nums, k))
}
