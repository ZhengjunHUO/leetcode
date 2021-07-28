package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	k := 0
	
	for _,v := range nums {
		if k < 1 || v > nums[k-1] {
			nums[k] = v
			k++
		}
	}

	fmt.Println(nums[:k])

	return k
}

func main() {
	nums := [][]int{[]int{1,1,2}, []int{0,0,1,1,1,2,2,3,3,4}}
	for _,v := range nums {
		fmt.Println(removeDuplicates(v))
	}
}
