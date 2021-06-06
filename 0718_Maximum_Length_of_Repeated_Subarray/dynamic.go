package main

import (
	"fmt"
)

func findLength(nums1 []int, nums2 []int) int {
	len1, len2, rslt := len(nums1), len(nums2), 0

	dp := make([][]int, len1+1)
	for x := range dp {
		dp[x] = make([]int, len2+1)
	}

	for i:=len1-1; i>=0; i-- {
		for j:=len2-1; j>=0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > rslt {
					rslt = dp[i][j]
				}
			}
		}
	}
	return rslt
}

func main() {
	nums1, nums2 := []int{1,2,3,2,1}, []int{3,2,1,4,7}
	fmt.Println(findLength(nums1, nums2))

	nums1, nums2 = []int{0,0,0,0,0}, []int{0,0,0,0,0}
	fmt.Println(findLength(nums1, nums2))
}
