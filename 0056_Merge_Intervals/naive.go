package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] } )

	rslt := [][]int{}
	rslt = append(rslt, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if rslt[len(rslt)-1][1] < intervals[i][0] {
			rslt = append(rslt, intervals[i])
		}else{
			if rslt[len(rslt)-1][1] < intervals[i][1] {
				rslt[len(rslt)-1][1] = intervals[i][1] 
			}
		}
	}
	
	return rslt
}

func main() {
	intervals := [][][]int{[][]int{[]int{1,3},[]int{2,6},[]int{8,10},[]int{15,18}}, [][]int{[]int{1,4}, []int{4,5}}}
	for i := range intervals {
		fmt.Println(merge(intervals[i]))
	}
}
