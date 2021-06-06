package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	h := make(map[int]int)

        for i := range nums {
		diff := target - nums[i]
		if elem, ok := h[diff]; ok {
			return []int{elem,i}
		}
				
		h[nums[i]] = i
	}

	fmt.Println("Violation of convention detected !")
	return []int{0}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}
