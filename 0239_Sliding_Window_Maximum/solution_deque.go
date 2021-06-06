package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums) 

	// 新元素在尾部加入，头部为当前范围内的最大值
	// 保存的值为index
	dq := godtype.NewDeque()
	rslt := make([]int, n-k+1)
	rsltIdx := 0    

	for i := range nums{
		// 去除已经落在范围外的局部最大值（检查最大值的index）
		for (!dq.IsEmpty()) && (dq.PeekFirst().(int) < i - k + 1) {
			dq.PopFirst()
		}	

		// 从尾向头去除小于当前nums[i]值的队列元素（检查index对应的值）
		for (!dq.IsEmpty()) && (nums[dq.PeekLast().(int)]) < nums[i] {
			dq.PopLast()
		}

		// 将当前元素的index插入队尾
		dq.PushLast(i)
		if i >= k - 1 {
			// 头部为局部最大值的index，将对应数值保存至结果
			rslt[rsltIdx] = nums[dq.PeekFirst().(int)]
			rsltIdx += 1
		}
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
