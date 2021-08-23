package main

import (
	"fmt"
)

/*
  思路是创建一些堆，只有比堆顶小的数可以放到堆上；如果数较大没有堆可以放就开新堆；
  如果有几个候选堆则选择最左边的；最后的堆数就是最大子序列
*/
func lengthOfLIS(nums []int) int {
	top, piles := make([]int, len(nums)), 0

	for i := range nums {
		l, r := 0, piles
		
		for l < r {
			m := l + (r - l)/2			
			// 取最左侧的堆放较小的数
			if top[m] >= nums[i] {
				r = m
			}else{
				l = m + 1
			}
		}

		// 数较大没有找到堆，开一个新堆
		if l == piles {
			piles++
		}
		// 当前数放在找到的堆顶
		top[l] = nums[i]
	}

	return piles
}

func main() {
	nums := [][]int{[]int{10,9,2,5,3,7,101,18}, []int{0,1,0,3,2,3}, []int{7,7,7,7,7,7,7}}	
	for i := range nums {
		fmt.Println(lengthOfLIS(nums[i]))
	}
}
