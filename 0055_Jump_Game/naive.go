package main

import (
	"fmt"
)

// 从目的地往前倒推
// 如果在位置i可以到达当前的目标，就可以更新目标为位置i，然后进入i-1继续; 如果不能到达，就继续往前看。
// 如果退出loop时target不是0意味着不能到达
func canJump(nums []int) bool {
	n := len(nums)
	targetIdx := n-1    

	for i:=n-2 ; i>=0; i-- {
		if nums[i] + i >= targetIdx {
			targetIdx = i
		}
	} 

	return targetIdx == 0
}

func main() {
	nums := []int{2,3,1,1,4}
	fmt.Println(canJump(nums))

	nums = []int{3,2,1,0,4}
	fmt.Println(canJump(nums))
}
