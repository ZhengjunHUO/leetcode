package main

import (
	"fmt"
)

func isSameAfterReversals(num int) bool {
	return num == 0 || num % 10 != 0
}

func main() {
	nums := []int{526, 1800, 0}
	for i := range nums {
		fmt.Println(isSameAfterReversals(nums[i]))
	}
}
