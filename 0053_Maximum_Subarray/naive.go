package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]

	for i:=1; i<len(nums); i++ {
		if temp:=sum+nums[i]; temp > nums[i] {
			sum = temp
		}else{
			// 如果cumsum加上当前元素还小于该元素，还不如从该元素开始
			sum = nums[i]
		}
	
		// 记录历史上出现过的最大cumsum
		if sum > max {
			max = sum
		} 
	}

	return max
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))

	nums = []int{1}
	fmt.Println(maxSubArray(nums))

	nums = []int{5,4,-1,7,8}
	fmt.Println(maxSubArray(nums))
}
