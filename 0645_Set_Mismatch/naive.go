package main

import (
	"fmt"
)

// 参考0041_First_Missing_Positive的思想
func findErrorNums(nums []int) []int {
	rslt := make([]int, 2)

	n := len(nums) + 1
	for i := range nums {
		nums[(nums[i]-1)%n] += n
	}

	for i := range nums {
		if nums[i]/n == 0 {
			rslt[1] = i + 1
		}else{
			if (nums[i]-n)/n > 0 {
				rslt[0] = i + 1
			}
		}
	}

	return rslt
}

func main() {
	nums := [][]int{[]int{1,2,2,4}, []int{1,1}, []int{4,3,2,7,8,5,3,1}}
	for i := range nums {
		fmt.Println(findErrorNums(nums[i]))
	}
}
