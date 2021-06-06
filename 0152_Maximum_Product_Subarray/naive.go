package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	rslt, cumMax, cumMin := nums[0], nums[0], nums[0] 
 
	for i:=1; i < len(nums); i++ {
		if nums[i] < 0 {
			cumMax, cumMin = cumMin, cumMax
		}

		if temp := cumMax * nums[i]; temp > nums[i] {
			cumMax = temp
		}else{
			cumMax = nums[i]
		}

		if temp := cumMin * nums[i]; temp < nums[i] {
			cumMin = temp
		}else{
			cumMin = nums[i]
		}

		if cumMax > rslt {
			rslt = cumMax
		}
	}
	
	return rslt
}

func main() {
	nums := []int{2,3,-2,4}
	fmt.Println(maxProduct(nums))

	nums = []int{-2,0,-1}
	fmt.Println(maxProduct(nums))
}
