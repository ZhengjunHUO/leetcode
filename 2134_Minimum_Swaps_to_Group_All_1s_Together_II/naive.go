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

func minSwaps(nums []int) int {
	var ones,localOnes,maxLocalOnes int 
	// 统计nums中1的数量
	for i := range nums {
		if nums[i] == 1 {
			ones++
		}
	}

	// 遍历两次nums模拟循环数列
	// 使用移动窗口，大小为1的数量
	for i:=0; i<2*len(nums)-1; i++ {
		// 前面的1掉出窗口范围则递减计数器
		if (i >= ones && nums[(i - ones)%len(nums)] == 1) {
			localOnes--
		}

		// 当前加入的元素为1则递增计数器
		if nums[i%len(nums)] == 1 {
			localOnes++
		}

		maxLocalOnes = max(maxLocalOnes, localOnes)
	}

	// 需要swap的次数：1的总数量减去窗口中最大出现的1的数量
	return ones - maxLocalOnes
}

func main() {
	nums := [][]int{[]int{0,1,0,1,1,0,0}, []int{0,1,1,1,0,0,1,1,0}, []int{1,1,0,0,1}}
	for i := range nums {
		fmt.Println("Rslt: ", minSwaps(nums[i]))
	}
}
