package main

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {
	// 时间复杂度为O(nlogn)
	sort.Ints(nums)	

	/*
		preVal: 前一个数正好等于n-1时其计算得到的最大值，用于衡量是否要取当前值n；如果前一个数<n-1则其值为0
		prepreVal: 值为n-1的数之前的历史最大值，使用rslt值更新
		currentVal: 值同为n的数的累计之和，用来和prepreVal相加后与preVal比较，较大的值被存在rslt中。
	*/

	prepreVal, preVal, currentVal, prevNum, rslt := 0, 0, nums[0], nums[0], nums[0]

	for i:=1; i<len(nums); i++ {
		if diff := nums[i] - prevNum; diff > 1 {
			prepreVal, preVal, currentVal, prevNum = rslt, 0, nums[i], nums[i]
		}else if diff == 1 {
		// 相差1
			prepreVal, preVal, currentVal, prevNum = preVal, rslt, nums[i], nums[i]
			
		}else{
		// 相等
			currentVal += nums[i]
		}	

		if temp := currentVal + prepreVal; temp > preVal {
			rslt = temp
		}else{
			rslt = preVal
		}

	}

	return rslt
}

func main() {
	nums := []int{3,4,2}
	fmt.Println(deleteAndEarn(nums))

	nums = []int{2,2,3,3,3,4}
	fmt.Println(deleteAndEarn(nums))
}
