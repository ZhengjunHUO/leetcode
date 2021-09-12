package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	rslt := []string{}
	if len(nums) == 0 {
		return rslt
	}

	prev, str := nums[0], fmt.Sprint(nums[0])
	
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] + 1 {
			continue
		}
	
		if nums[i-1] != prev {
			str = str + "->" + fmt.Sprint(nums[i-1])
		}

		rslt = append(rslt, str)
		prev, str = nums[i], fmt.Sprint(nums[i])
	}	

	
	if nums[len(nums)-1] != prev {
		str = str + "->" + fmt.Sprint(nums[len(nums)-1])
	}
	rslt = append(rslt, str)

	return rslt
}

func main() {
	nums := [][]int{[]int{0,1,2,4,5,7}, []int{0,2,3,4,6,8,9}, []int{}, []int{-1}, []int{0}}
	for i := range nums {
		fmt.Println(summaryRanges(nums[i]))
	}
}
