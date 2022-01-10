package main

import (
	"fmt"
	"sort"
)

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func FindLeftBound(slice [][]int, target int, less func(i, target int) bool) int {
	n := len(slice)
	l, r := 0, n

	for l < r {
		m := l + (r-l)/2
		if less(m, target) {
			l = m + 1
		}else{
			r = m
		}
	}

	return l
}

func maxValue(events [][]int, k int) int {
	// 准备一个n+1乘以k+1的表，第n+1行和第0列初值为0, 其余值为-1
	// 从右上角[0,k]出发(即处于第一个event,且余额为k时)向下调用递归函数
	// 再自底向上计算填充dp表
	dp := make([][]int, len(events)+1)	
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			if j == 0 || i == len(events) {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = -1
		}
	}

	// 按开始时间升序排序
	sort.SliceStable(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}

		return events[i][0] < events[j][0]
	})

	// 返回右上角的值
	return recursive(&events, &dp, 0, k)
}

func recursive(events *[][]int, dp *[][]int, pos, k int) int {
	// 查表，避免重复计算
	if (*dp)[pos][k] != -1 {
		return (*dp)[pos][k]
	}

	// 计算紧接着的下一个事件的index
	// 对起始时间使用二分搜索找到"值为当前事件结束时间+1"的左边界
	/*
	var idx int
	for idx = pos + 1; idx < len(*events); idx++ {
		if (*events)[idx][0] > (*events)[pos][1] {
			break
		}
	}
	*/

	idx := FindLeftBound((*events)[pos+1:], (*events)[pos][1]+1, func(i, target int) bool {
		return ((*events)[pos+1:])[i][0] < target
	}) + (pos + 1)
	//fmt.Printf("Get index: %d\n", idx)

	// 1) 不取当前事件(此时k也没有消耗)，其值取下一个事件和k计算所得
	// 2) 取当前事件，其值为"当前事件的价值"加上"下一个可取事件和k-1"的计算所得
	(*dp)[pos][k] = max(recursive(events, dp, pos+1, k), recursive(events, dp, idx, k-1) + (*events)[pos][2])
	return (*dp)[pos][k]
}

func main() {
	events := [][][]int{[][]int{[]int{1,2,4},[]int{3,4,3},[]int{2,3,1}}, [][]int{[]int{1,2,4},[]int{3,4,3},[]int{2,3,10}}, [][]int{[]int{1,1,1},[]int{2,2,2},[]int{3,3,3},[]int{4,4,4}}}
	k := []int{2,2,3}

	for i := range events {
		fmt.Println(maxValue(events[i], k[i]))
	}
}
