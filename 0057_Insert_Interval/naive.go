package main

import (
	"fmt"
)

// 也可以遍历intervals, 判断并决定是否直接append到结果还是和newInterval融合: O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	begin, end := binarySearch(intervals, newInterval, 0), binarySearch(intervals, newInterval, 1)

	if begin > 0 && newInterval[0] <= intervals[begin-1][1] {
		begin--
	}
	if end == len(intervals) || newInterval[1] < intervals[end][0] {
		end--
	}

	fmt.Println("begin, end: ", begin, end)

	if intervals[begin][0] < newInterval[0]{
		newInterval[0] = intervals[begin][0]
	}

	if intervals[end][1] > newInterval[1] {
		newInterval[1] = intervals[end][1]
	}

	fmt.Println("begin, end: ", begin, end)
	rslt := append(intervals[:begin], newInterval)
	rslt = append(rslt, intervals[end+1:]...)

	return rslt
}

func binarySearch(intervals [][]int, newInterval []int, index int) int {
	l, r := 0, len(intervals) - 1

	for l <= r {
		m := l + (r-l)/2
		if newInterval[index] < intervals[m][index] {
			r = m - 1
		}else if newInterval[index] > intervals[m][index] {
			l = m + 1
		}else{
			return m
		}
	}
	
	return l
}

func main() {
	intervals := [][][]int{[][]int{[]int{1,3}, []int{6,9}}, [][]int{[]int{1,2},[]int{3,5},[]int{6,7},[]int{8,10},[]int{12,16}}, [][]int{}, [][]int{[]int{1,5}}, [][]int{[]int{1,5}}}
	newIntervals := [][]int{[]int{2,5}, []int{4,8}, []int{5,7}, []int{2,3}, []int{2,7}}

	for i := range intervals {
		fmt.Println(insert(intervals[i], newIntervals[i]))
	}
}
