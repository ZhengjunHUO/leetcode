package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
} 

//类似1024_Video_Stitching
func jump(nums []int) int {
	cnt, currend, maxend := 0, 0, 0

	for i := range nums {
		// 在currend的范围里，计算最远可以到达什么位置
		maxend = max(maxend, i+nums[i]) 

		// 上一步能覆盖的范围到头了，需要增加步数
		// 新的一步的范围为上一步中找到的最远的位置
		if i == currend {
			cnt++
			currend = maxend

			if currend >= len(nums)-1 {
				return cnt
			}
		}
	}

	return cnt
}

func main() {
	nums := []int{2,3,1,1,4}
	fmt.Println(jump(nums))

	nums = []int{2,3,0,1,4}
	fmt.Println(jump(nums))
}
