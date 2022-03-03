package main

import (
	"fmt"
)

const MAXINT64 = int64(^uint64(0) >> 1)

func countTrips(time []int, givenTime int64) int {
	var ret int
	for i := range time {
		ret += int(givenTime/int64(time[i]))
	}
	return ret
}

// 将问题转换为二分搜索，countTrips为目标时间的递增函数，搜索其左边界找其最小值
func minimumTime(time []int, totalTrips int) int64 {
	l, r := int64(0), MAXINT64

	for l < r {
		m := l + (r-l) / 2
		if countTrips(time, m) >= totalTrips {
			r = m
		}else{
			l = m + 1
		}
	}

	return l
}

func main() {
	time := [][]int{[]int{1,2,3}, []int{2}}
	totalTrips := []int{5, 1}

	for i := range time {
		fmt.Println(minimumTime(time[i], totalTrips[i]))
	}
}
