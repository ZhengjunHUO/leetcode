package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	maj, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			maj, count = nums[i], 1
			continue
		}

		if maj == nums[i] {
			count++
		}else{
			count--
		}

	}

	return maj	
}

func main() {
	nums := [][]int{[]int{3,2,3}, []int{2,2,1,1,1,2,2}, []int{3,1,2,1,1,3,1}}
	for i := range nums {
		fmt.Println(majorityElement(nums[i]))
	}
}
