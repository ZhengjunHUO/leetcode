package main

import (
	"fmt"
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	for i := 0; i < len(intervals) - 1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}

	return true
}

func main() {
	intervals := [][][]int{[][]int{[]int{0,30},[]int{5,10},[]int{15,20}}, [][]int{[]int{7,10},[]int{2,4}}}
	for i := range intervals {
		fmt.Println(canAttendMeetings(intervals[i]))
	}
}
