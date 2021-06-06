package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i:=0; i<n-1; i++ {
		lp, rp, goal := i+1, n-1, target - numbers[i] 
		for lp<=rp {
			mid := lp + (rp-lp)/2
			switch {
			case numbers[mid]==goal:
				return []int{i+1, mid+1}
			case numbers[mid]>goal:
				rp = mid-1
			case numbers[mid]<goal:
				lp = mid+1
			}
		}
	}    

	fmt.Println("Error in problem")
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15},9))
	fmt.Println(twoSum([]int{2,3,4},6))
	fmt.Println(twoSum([]int{-1,0},-1))
	fmt.Println(twoSum([]int{1,4,7,9,15,18,22},25))
	fmt.Println(twoSum([]int{1,4,7,9,15,18,22},11))
}
