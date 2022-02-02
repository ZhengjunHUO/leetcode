package main

import (
	"fmt"
)

func nextPositive(nums []int, pos *int) {
	*pos += 1
	for *pos < len(nums) && nums[*pos] < 0 {
		*pos += 1
	}
}

func nextNegative(nums []int, pos *int) {
	*pos += 1
	for *pos < len(nums) && nums[*pos] > 0 {
		*pos += 1
	}

}

func rearrangeArray(nums []int) []int {
	pPos, pNeg := -1, -1
	nextPositive(nums, &pPos)
	nextNegative(nums, &pNeg)

	rslt := make([]int, 0, len(nums))
	for pPos < len(nums) {
		rslt = append(rslt, nums[pPos])
		nextPositive(nums, &pPos)
		rslt = append(rslt, nums[pNeg])
		nextNegative(nums, &pNeg)
	}

	return rslt
}

func main() {
	nums := [][]int{[]int{3,1,-2,-5,2,-4}, []int{-1,1}}
	for i := range nums {
		fmt.Println(rearrangeArray(nums[i]))
	}
}
