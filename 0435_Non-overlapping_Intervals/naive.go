package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	cnt := 0

	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] } )
	curr := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[curr][1] > intervals[i][0] {
			cnt++
		}else{
			curr++
		}
	}	

	return cnt
}

func main() {
	intervals := [][][]int{[][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{1,3}}, [][]int{[]int{1,2},[]int{1,2},[]int{1,2}}, [][]int{[]int{1,2},[]int{2,3}}}
	for i := range intervals {
		fmt.Println(eraseOverlapIntervals(intervals[i]))
	}
}
