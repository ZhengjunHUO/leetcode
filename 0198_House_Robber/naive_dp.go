package main

import (
	"fmt"
)

/*
rslt[i] = max(rslt[i-2]+nums[i], rslt[i-1])
比较抢夺和不抢夺当前房子的收益取最大值
*/

func rob(nums []int) int {
	n := len(nums)
	// 只需保存和刷新rslt[i-2]和rslt[i-1]两个值即可
	preprev, prev := 0, nums[0]
	
	for i:=1; i<n; i++ {
		if current := preprev + nums[i]; current > prev {
			preprev, prev = prev, current
		}else{
			preprev, prev = prev, prev
		}
		
	}

	return prev
}

func main() {
	nums := []int{1,2,3,1}
	fmt.Println(rob(nums))

	nums = []int{2,7,9,3,1}
	fmt.Println(rob(nums))

	nums = []int{3,2,1,4}
	fmt.Println(rob(nums))
}
