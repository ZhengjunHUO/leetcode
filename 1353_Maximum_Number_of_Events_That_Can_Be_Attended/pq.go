package main

import (
	"fmt"
	"sort"

	"github.com/ZhengjunHUO/godtype"
)

func maxEvents(events [][]int) int {
	// 按开始日期升序排序
	sort.SliceStable(events, func(i, j int) bool{
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}

		return events[i][0] < events[j][0]
	})

	// 声明一个min PQ
	pq := godtype.NewPQ([]int{}, []int{}, true)

	i, n, start, rslt := 0, len(events), 0, 0
	for pq.Size() > 0 || i < n {
		// 记录"第一个"活动开始时间start
		if pq.Size() == 0 {
			start = events[i][0]
		}

		// 加入开始时间早于start的事件(关心其结束时间), pq保证结束时间最早的事件排在最前
		for i < n && events[i][0] <= start {
			pq.Push(events[i][1], events[i][1])
			i++
		}

		// 参加最早结束的事件
		pq.Pop()
		// 占用了一天
		start += 1
		// 结果加一
		rslt += 1

		// 放弃已经过期的事件
		for pq.Size() > 0 && pq.Peek().(int) < start {
			pq.Pop()
		}
	}

	return rslt
}

func main() {
	events := [][][]int{[][]int{[]int{1,2},[]int{2,3},[]int{3,4}}, [][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{1,2}}}
	for i := range events {
		fmt.Println(maxEvents(events[i]))
	}
}
