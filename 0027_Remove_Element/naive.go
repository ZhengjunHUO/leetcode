package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	k := 0

	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}

	fmt.Println(nums[:k])
	return k
}

func main() {
	nums := [][]int{[]int{3,2,2,3}, []int{0,1,2,2,3,0,4,2}}
	vals := []int{3, 2}

	for i := range nums {
		fmt.Println(removeElement(nums[i], vals[i]))
	}
}
