package main

import (
	"fmt"
	"sort"
)

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func maxTwoEvents(events [][]int) int {
	// 只需要挑选两个interval，所以maxSince只需要记录
	// 到目前为止结束的价值最大的interval的值
	rslt, maxSince := 0, 0
	slice := make([][]int, 0, len(events))
	
	/*
	  三元组为
		起始slice {起始时间，1，权重}
		结束slice {结束时间，0，权重}
	*/	

	for i := range events {
		slice = append(slice, []int{events[i][0], 1, events[i][2]})
		slice = append(slice, []int{events[i][1]+1, 0, events[i][2]})
	}

	// 按时间升序排序，如果相同则结束slice排在前面
	sort.SliceStable(slice, func(i, j int) bool{
		if slice[i][0] == slice[j][0] {
			return slice[i][1] < slice[j][1]
		}
		
		return slice[i][0] < slice[j][0]
	})

	for i := range slice {
		// 找到一个起始slice，其值加上maxSince如果较大则更新结果
		if slice[i][1] == 1 {
			rslt = max(rslt, maxSince + slice[i][2])
		// 找到一个结束slice, 其值如果大于maxSince则更新之
		}else{
			maxSince = max(maxSince, slice[i][2])
		}
	}

	return rslt
}

func main() {
	events := [][][]int{[][]int{[]int{1,3,2},[]int{4,5,2},[]int{2,4,3}}, [][]int{[]int{1,3,2},[]int{4,5,2},[]int{1,5,5}},[][]int{[]int{1,5,3},[]int{1,5,1},[]int{6,6,5}}}
	for i := range events {
		fmt.Println(maxTwoEvents(events[i]))
	}
}
