package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ret := 100001
	for i:=0; i<len(nums)-k+1; i++ {
		ret = min(ret, nums[i+k-1]-nums[i])
	}

	return ret
}

func main() {
	nums := [][]int{[]int{90}, []int{9,4,1,7}}
	k := []int{1, 2}

	for i := range nums {
		fmt.Println(minimumDifference(nums[i], k[i]))
	}
}
