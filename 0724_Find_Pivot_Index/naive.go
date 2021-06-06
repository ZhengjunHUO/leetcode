package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	leftSum,prev := 0,0
	rightSum := 0
	for _,v := range nums {
		rightSum += v
	}    

	for i,v := range nums {
		leftSum += prev
		prev = v
		rightSum -= v	
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{1,7,3,6,5,6}
	fmt.Println(pivotIndex(nums))

	nums = []int{1,2,3}
	fmt.Println(pivotIndex(nums))

	nums = []int{2,1,-1}
	fmt.Println(pivotIndex(nums))
}
