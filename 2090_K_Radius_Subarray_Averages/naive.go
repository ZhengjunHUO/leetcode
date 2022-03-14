package main

import (
	"fmt"
)

func getAverages(nums []int, k int) []int {
	n := len(nums)
	ret,acc,radius := make([]int, n), make([]int, n+1), 2*k+1

	for i := range nums {
		acc[i+1] = acc[i]+nums[i]
		ret[i] = -1
	}

	for j := k; j <= n-k-1; j++ {
		ret[j] = (acc[j+k+1]-acc[j-k])/radius
	}

	return ret
}

func main() {
	nums := [][]int{[]int{7,4,3,9,1,8,5,2,6}, []int{100000}, []int{8}}
	k := []int{3,0,100000}

	for i := range nums {
		fmt.Println(getAverages(nums[i], k[i]))
	}
}
