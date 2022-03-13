package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
  根据hint提示，每一个种子开花的时间为之前所有种子的准备时间之和加上当前种子生长时间
  参考@rh137的解法，按花所需要的生长时间降序排列，使用贪心算法
*/
func earliestFullBloom(plantTime []int, growTime []int) int {
	// 将两个数列合并并排序 
	zip := make([][2]int, len(plantTime))
	for i := range zip{
		zip[i][0], zip[i][1] = plantTime[i], growTime[i]
	}
	sort.SliceStable(zip, func (a, b int) bool {
		return zip[a][1] > zip[b][1]
	})

	// start为"之前所有种子的准备时间之和"
	start, ret := 0, 0
	for i := range zip {
		ret = max(ret, start+zip[i][0]+zip[i][1])
		// 为下一轮更新start
		start += zip[i][0]
	} 

	return ret
}

func main() {
	pt := [][]int{[]int{1,4,3}, []int{1,2,3,2}, []int{1}}
	gt := [][]int{[]int{2,3,1}, []int{2,1,2,1}, []int{1}}
	for i := range pt {
		fmt.Println(earliestFullBloom(pt[i], gt[i]))
	}
}
