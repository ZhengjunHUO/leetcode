package main

import (
	"fmt"
)

func sortColors(nums []int)  {
	l, m, r := 0, 0, len(nums) - 1

	for m<=r {
		switch nums[m] {
		case 0:
			nums[l], nums[m] = nums[m], nums[l]
			l++
			m++
		case 1:
			m++
		case 2:
			nums[m], nums[r] = nums[r], nums[m]
			r--
		}
	}
}

func main() {
	nums := [][]int{[]int{2,0,2,1,1,0}, []int{2,0,1}, []int{0}, []int{1}, []int{1,2,2,0,1,0,2,0,1}}
	for i := range nums {
		sortColors(nums[i])
		fmt.Println(nums[i])
	}
}
