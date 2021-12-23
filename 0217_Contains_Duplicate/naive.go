package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	dict := make(map[int]bool)
	
	for _,v := range nums {
		if dict[v] == true {
			return true
		}

		dict[v] = true
	}

	return false
}

func main() {
	nums := [][]int{[]int{1,2,3,1}, {1,2,3,4}, {1,1,1,3,3,4,3,2,4,2}}
	for i := range nums {
		fmt.Println(containsDuplicate(nums[i]))
	}	
}
