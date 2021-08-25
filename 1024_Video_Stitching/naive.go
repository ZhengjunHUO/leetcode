package main

import (
	"fmt"
	"sort"
)

func max (a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 与0045_Jump_Game_II类似
func videoStitching(clips [][]int, time int) int {
	// 按起始时间升序排列，如相同按终止时间降序排列
	sort.SliceStable(clips, func (i,j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1] 
		}
		
		return clips[i][0] < clips[j][0]
	})

	cnt, currend, maxend, i := 0, 0, 0, 0

	for i < len(clips) && clips[i][0] <= currend {
		// 所有起始坐标小于当前终止坐标的区间中，寻找能涵盖最远的距离
		for i < len(clips) && clips[i][0] <= currend {
			maxend = max(maxend, clips[i][1])
			i++
		}
		
		// 上个for中选中了一个区间，cnt加1
		cnt++
		// 将当前终止坐标更新为该区间坐标
		currend = maxend
		if currend >= time {
			return cnt
		}
	}

	return -1
}

func main() {
	clips := [][][]int{[][]int{[]int{0,2},[]int{4,6},[]int{8,10},[]int{1,9},[]int{1,5},[]int{5,9}}, [][]int{[]int{0,1},[]int{1,2}}, [][]int{[]int{0,1},[]int{6,8},[]int{0,2},[]int{5,6},[]int{0,4},[]int{0,3},[]int{6,7},[]int{1,3},[]int{4,7},[]int{1,4},[]int{2,5},[]int{2,6},[]int{3,4},[]int{4,5},[]int{5,7},[]int{6,9}}, [][]int{[]int{0,4},[]int{2,8}}}
	times := []int{10, 5, 9, 5}

	for i := range clips {
		fmt.Println(videoStitching(clips[i], times[i]))
	}
}
