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

func minimumRemoval(beans []int) int64 {
	// 排序
	sort.Ints(beans)

	// 求总豆子数
	var sum int
	for i := range beans {
		sum += beans[i]
	}

	ret := sum
	// 如果以i为基准，i之前的数字需要清零，i后的数字都需要减少到beans[i]
	// 需要移除的数量为：总数减去剩下的数量
	for i := range beans {
		ret = min(ret, sum - (len(beans) - i)*beans[i])
	}

	return int64(ret)
}

func main() {
	beans := [][]int{[]int{4,1,6,5}, []int{2,10,3,2}}
	for i := range beans {
		fmt.Println(minimumRemoval(beans[i]))
	}
}
