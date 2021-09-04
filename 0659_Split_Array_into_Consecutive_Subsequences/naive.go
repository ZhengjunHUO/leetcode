package main

import (
	"fmt"
)

func isPossible(nums []int) bool {
	count, wait := make(map[int]int), make(map[int]int) 

	for i := range nums {
		count[nums[i]]++
	}

	for i := range nums {
		if count[nums[i]] == 0 {
			continue
		}

		if wait[nums[i]] > 0 {
			count[nums[i]]--
			wait[nums[i]]--
			wait[nums[i]+1]++	
		}else if count[nums[i]+1] > 0 && count[nums[i]+2] > 0 {
			count[nums[i]]--
			count[nums[i]+1]--
			count[nums[i]+2]--
			wait[nums[i]+3]++	
		}else{
			return false
		}
	}

	return true
}

func main() {
	nums := [][]int{[]int{1,2,3,3,4,5}, []int{1,2,3,3,4,4,5,5}, []int{1,2,3,4,4,5}}
	for i := range nums {
		fmt.Println(isPossible(nums[i]))
	}
}
