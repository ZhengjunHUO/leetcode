package main

import (
	"fmt"
)

// 使用backtracking
func most(questions [][]int, idx int, curr int, max *int) {
	// 终止条件
	if idx >= len(questions) {
		// 更新max
		if curr > *max {
			*max = curr
		}
		return
	}

	// 不取当前点
	most(questions, idx+1, curr, max)
	// 取当前点
	most(questions, idx+questions[idx][1]+1, curr+questions[idx][0], max)
}

func mostPoints(questions [][]int) int64 {
	var max int	
	most(questions, 0, 0, &max)
	return int64(max)
}

func main() {
	qs := [][][]int{[][]int{[]int{3,2},[]int{4,3},[]int{4,4},[]int{2,5}}, [][]int{[]int{1,1},[]int{2,2},[]int{3,3},[]int{4,4},[]int{5,5}}}
	for i := range qs {
		fmt.Println(mostPoints(qs[i]))
	}
}
