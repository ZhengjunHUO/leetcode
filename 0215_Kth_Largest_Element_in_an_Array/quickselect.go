package main

import (
	"fmt"
)

/* 
从选择最左边的值为pivot开始，将比它小的值放在其左侧，大的值放在右侧（完成部分排序）
不断缩小区域找到解，故理论时间复杂度为O(nlogn)，实际平均可以达到O(n)
最差的情况为O(n^2)，如升序数列从左边开始找最大值的情况
*/
func findKthLargest(nums []int, k int) int {
	// 目标index
	target := len(nums) - k
	l, r := 0, len(nums) - 1
	var rslt int
	
	loop: for l <= r {
		p := l

		// 目标是比nums[l]小的元素在本轮结束后都排列在其左边
		for i:=l+1; i<=r; i++ {
			if nums[i] < nums[l] {
				p++
				// 遇到较小的值将其换到前面，确保p之前的值小于标杆值
				nums[i], nums[p] = nums[p], nums[i]
			}
		}
		// 遍历完后swap p指向的值和标杆值(此时仅可以确定标杆值完成了排序)
		nums[l], nums[p] = nums[p], nums[l]

		switch {
			// 如果目标坐落在pivot的右侧，则继续搜寻右侧区域
			case p < target:
				l = p + 1
			// 如果目标坐落在pivot的左侧，则继续搜寻左侧区域
			case p > target:
				r = p - 1
			// 找到目标
			case p == target:
				rslt = nums[p]
				break loop
		}
	}

	return rslt
}

func main() {
	fmt.Println(findKthLargest([]int{3,2,1,5,6,4}, 2))
	fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
}
