package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	n := len(nums)
	switch {
	case n == 1:
		return nums[0]
	case n == 2:
		if nums[0] >= nums[1] {
			return nums[0]
		}else{
			return nums[1]
		}
	default:
		m1, m2, m3 := -2147483649, -2147483649, -2147483649
		for _, v := range nums {
			if v <= m3 || v == m2 || v == m1 {
				continue
			}

			m3 = v
			if m3 > m2 {
				m2, m3 = m3, m2
			}
			if m2 > m1 {
				m1, m2 = m2, m1
			}
		}
		return m3
	}
}

func main() {
	ints := [][]int{[]int{3,2,1}, []int{1,2}, []int{2,2,3,1}}

	for i := range ints {
		fmt.Println(thirdMax(ints[i]))
	}
}
