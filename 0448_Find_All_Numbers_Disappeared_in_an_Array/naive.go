package main

import (
	"fmt"
)

// 参考0041_First_Missing_Positive的思想
func findDisappearedNumbers(nums []int) []int {
	// 值是从[1,n]的，所以需要+1
	n := len(nums) + 1
	for i := range nums {
		// 同理此处需要减1
		nums[(nums[i]%n)-1] += n
	}

	rslt := []int{}
	for i := range nums {
		// 如果index没有对应的值出现过其值没有加过n故<n
		if nums[i]/n == 0 {
			rslt = append(rslt, i+1)
		}
	}
	
	return rslt
}

func main() {
	nums := [][]int{[]int{4,3,2,7,8,2,3,1}, []int{1,1}}
	for i := range nums {
		fmt.Println(findDisappearedNumbers(nums[i]))
	}
}
