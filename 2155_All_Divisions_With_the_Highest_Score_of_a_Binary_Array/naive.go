package main

import (
	"fmt"
)

func maxScoreIndices(nums []int) []int {
	count := make([]int, len(nums)+1)
	ret := []int{}
	left, right, max := 0, 0, 0

	// 1st pass from left to right, count zero's number
	for i := range count {
		count[i] += left

		if i < len(nums) && nums[i] == 0 {
			left++
		}
	}

	// 2nd pass from right to left, count one's number and add to left part's sum
	for i := len(count)-1; i >= 0; i-- {
		count[i] += right

		// record max and update associated index
		if count[i] > max {
			max = count[i]
			ret = []int{i}
		}else if count[i] == max {
			ret = append(ret, i)
		}

		if i >= 1 && nums[i-1] == 1 {
			right++
		}
	}

	return ret
}

func main() {
	nums := [][]int{[]int{0,0,1,0}, []int{0,0,0}, []int{1,1}}
	for i := range nums {
		fmt.Println(maxScoreIndices(nums[i]))
	}
}
