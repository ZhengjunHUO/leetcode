package main

import (
	"fmt"
)

func countKDifference(nums []int, k int) int {
	dict := make(map[int]int)
	var ret int

	for i := range nums {
		if v, ok := dict[nums[i]+k]; ok {
			ret += v
		}

		if v, ok := dict[nums[i]-k]; ok {
			ret += v
		}

		if _, ok := dict[nums[i]]; ok {
			dict[nums[i]] += 1
		}else{
			dict[nums[i]] = 1
		}
	}

	return ret
}

func main() {
	nums := [][]int{[]int{1,2,2,1}, []int{1,3}, []int{3,2,1,5,4}}
	k := []int{1,3,2}

	for i := range nums {
		fmt.Println(countKDifference(nums[i], k[i]))
	}
}
