package main

import (
	"fmt"
	"sort"
)

// 思路同0435_Non-overlapping_Intervals/greedy.go即按结束坐标升序排列的贪心算法
func findMinArrowShots(points [][]int) int {
	sort.SliceStable(points, func (i, j int) bool { return points[i][1] < points[j][1]})
	
	cnt, currend := 0, - (1 << 31)
	for i := range points {
		if points[i][0] > currend {
			cnt++
			currend = points[i][1]
		}
	}

	return cnt
}

func main() {
	ps := [][][]int{[][]int{[]int{10,16},[]int{2,8},[]int{1,6},[]int{7,12}}, [][]int{[]int{1,2},[]int{3,4},[]int{5,6},[]int{7,8}}, [][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{4,5}}}
	for i := range ps {
		fmt.Println(findMinArrowShots(ps[i]))
	}
}
