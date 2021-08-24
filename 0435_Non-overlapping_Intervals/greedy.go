package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	// 按结束坐标升序排列
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] } )
	
	// 保留区间数
	cnt := 0
	currend := -50000

	for i := range intervals {
		// 如果当前区间的开始坐标大于等于上一个选中保留区间的末尾坐标表示不重合
		// 递增计数器，用当前区间的末尾坐标来评判后续区间
		if intervals[i][0] >= currend {
			cnt++
			currend = intervals[i][1]
		}
	}	

	return len(intervals) - cnt
}

func main() {
	intervals := [][][]int{[][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{1,3}}, [][]int{[]int{1,2},[]int{1,2},[]int{1,2}}, [][]int{[]int{1,2},[]int{2,3}}}
	for i := range intervals {
		fmt.Println(eraseOverlapIntervals(intervals[i]))
	}
}
