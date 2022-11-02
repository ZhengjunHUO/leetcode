package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type Node struct {
	Index	int
	Value	int
}

// 同0239_Sliding_Window_Maximum一样需要使用deque
func maxResult(nums []int, k int) int {
	dq := datastruct.NewDeque([]Node{})
	curmax := 0

	for i:=0; i<len(nums); i++ {
		if dq.IsEmpty() {
			curmax = nums[i]
		}else{
			curmax = nums[i] + dq.PeekFirst().Value
		}

		for !dq.IsEmpty() && dq.PeekLast().Value < curmax {
			dq.PopLast()
		}

		dq.PushLast(Node{i, curmax})

		if dq.PeekFirst().Index <= i - k {
			dq.PopFirst()
		}
	}

	return curmax
}

func main() {
	nums := []int{1,-1,-2,4,-7,3}
	k := 2
	fmt.Println(maxResult(nums, k))

	nums = []int{10,-5,-2,4,0,3}
	k = 3
	fmt.Println(maxResult(nums, k))

	nums = []int{1,-5,-20,4,-1,3,-6,-3}
	k = 2
	fmt.Println(maxResult(nums, k))

	// should be 6
	nums = []int{1,-1,-2,-7,4,3}
	k = 2
	fmt.Println(maxResult(nums, k))

	// should be 22
	nums = []int{1,-1,-2,4,-7,3,10,-5,-2,4,0,3}
	k = 2
	fmt.Println(maxResult(nums, k))

	// should be 25
	k = 3
	fmt.Println(maxResult(nums, k))

	// should be 3
	nums = []int{1,-1,-2,-3,-4,-5,-7,4,3}
	k = 3
	fmt.Println(maxResult(nums, k))
}
