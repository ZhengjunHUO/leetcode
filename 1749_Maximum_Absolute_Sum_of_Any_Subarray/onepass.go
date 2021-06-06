package main

import (
	"fmt"
)

func maxAbsoluteSum(nums []int) int {
	var min, max, cum int

	for _,v := range nums {
		cum += v

		if cum < min {
			min = cum
		}
		if cum > max {
			max = cum
		}
	}

	return max - min
}

func main() {
	nums := []int{1,-3,2,3,-4}
	fmt.Println(maxAbsoluteSum(nums))

	nums = []int{2,-5,1,-4,3,-2}
	fmt.Println(maxAbsoluteSum(nums))
}
