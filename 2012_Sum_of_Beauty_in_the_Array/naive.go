package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func sumOfBeauties(nums []int) int {
	n := len(nums)
	pref, suff := make([]int, n), make([]int, n)

	pref[0] = nums[0]
	for i:=1; i<n; i++ {
		pref[i] = max(pref[i-1], nums[i])
	}

	suff[n-1] = nums[n-1]
	for i:=n-2; i>=0; i-- {
		suff[i] = min(suff[i+1], nums[i])
	}

	fmt.Println(pref)
	fmt.Println(suff)

	var ret int
	for i:=1; i<=n-2; i++ {
		if nums[i]>pref[i-1] && nums[i]<suff[i+1] {
			ret += 2
			continue
		}

		if nums[i]>nums[i-1] && nums[i]<nums[i+1] {
			ret += 1
		}
	}

	return ret
}

func main() {
	nums := [][]int{[]int{1,2,3}, []int{2,4,6,4}, []int{3,2,1}}
	for i := range nums {
		fmt.Println(sumOfBeauties(nums[i]))
	}
}
