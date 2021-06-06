package main

import (
	"fmt"
)

func maxAbsoluteSum(nums []int) int {
	lp, rp, max, cum, temp := 0, 0, 0, 0, 0

	for rp < len(nums) {
		cum += nums[rp]
		rp += 1

		if cum < 0 {
			temp = -cum
		}else{
			temp = cum
		}

		if temp >= max {
			max = temp
		}else{
			for lp < rp - 1 {
				cum -= nums[lp]
				lp += 1

				if cum < 0 {
					temp = -cum
				}else{
					temp = cum
				}

				if temp > max {
					max = temp
					break
				}
			} 
		}
	}

	return max
}

func main() {
	nums := []int{1,-3,2,3,-4}
	fmt.Println(maxAbsoluteSum(nums))

	nums = []int{2,-5,1,-4,3,-2}
	fmt.Println(maxAbsoluteSum(nums))
}
