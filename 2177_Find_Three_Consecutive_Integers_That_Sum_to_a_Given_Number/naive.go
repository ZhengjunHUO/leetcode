package main

import (
	"fmt"
)

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}

	n := num/3
	return []int64{n-1, n, n+1}
}

func main() {
	nums := []int64{33, 4}
	for i := range nums {
		fmt.Println(sumOfThree(nums[i]))
	}
}
