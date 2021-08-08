package main

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	// 按起点升序排列，如起点相等则按终点降序排列
	sort.SliceStable(intervals, func(i, j int) bool { 
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0] 
	})

	// 只需要关心右边界
	rslt, r := 0, -1

	for i := range intervals {
		// 部分重叠或不相交的情况，rslt需要+1
		if intervals[i][1] > r {
			r = intervals[i][1]
			rslt++
		}
	}
	
	return rslt
}

func main() {
	itvs := [][][]int{[][]int{[]int{1,4},[]int{3,6},[]int{2,8}}, [][]int{[]int{1,4},[]int{2,3}}, [][]int{[]int{0,10},[]int{5,12}}, [][]int{[]int{3,10},[]int{4,10},[]int{5,11}}, [][]int{[]int{1,2},[]int{1,4},[]int{3,4}}}
	for _,v := range itvs {
		fmt.Println(removeCoveredIntervals(v))
	}
}
