package main

import (
	"fmt"
)

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n, current, max := len(nums), 0, 0
	itvsum, lmaxidx, rmaxidx := make([]int, n-k+1), make([]int, n-k+1), make([]int, n-k+1)
	rslt1, rslt2, rslt3 := -1, -1, -1

	// 计算从位置i开始长度为k的各个interval的sum
	for i:=0; i<n; i++ {
		current += nums[i]	
		if (i>=k) {
			current -= nums[i-k]	
		}
		if (i>=k-1) {
			itvsum[i-k+1] = current
		}
	}

	// 从左到右找到每个时刻的历史最大值
	for i:=0; i<len(lmaxidx); i++ {
		if itvsum[i] > itvsum[max] {
			max = i
		}
		lmaxidx[i] = max
	}
	
	// 从右到左找到每个时刻的历史最大值，注意此处判断条件为>=
	max = len(rmaxidx) - 1
	for i:=len(rmaxidx) - 1; i>=0; i-- {
		if itvsum[i] >= itvsum[max] {
			max = i
		}
		rmaxidx[i] = max
	}
    
	for i:=k; i<len(itvsum)-k; i++ {
		l, r := lmaxidx[i-k], rmaxidx[i+k]
		if rslt1 == -1 || itvsum[l] + itvsum[i] + itvsum[r] > itvsum[rslt1] + itvsum[rslt2] + itvsum[rslt3] {
			rslt1, rslt2, rslt3 = l, i, r
		}
	}
	
	return []int{rslt1, rslt2, rslt3}
}

func main() {
	nums := []int{1,2,1,2,6,7,5,1}
	k := 2
	fmt.Println(maxSumOfThreeSubarrays(nums, k))	

	nums = []int{1,2,1,2,1,2,1,2,1}
	k = 2
	fmt.Println(maxSumOfThreeSubarrays(nums, k))	
}
