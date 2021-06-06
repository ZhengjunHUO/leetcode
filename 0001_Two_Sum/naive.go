package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	base := make([]int, len(nums))   
	for i := range base {
		base[i] = target - nums[i]
	}

        for j:= range base {
		for k := range nums {
			if base[j] == nums[k] && j !=k {
				return []int{j,k}
			}
		} 
	}

	fmt.Println("Violation of convention detected !")
	return []int{0}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}
