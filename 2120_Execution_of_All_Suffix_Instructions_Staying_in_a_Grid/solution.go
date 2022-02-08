package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 参考@hxu10的O(len(s))方法
func executeInstructions(n int, startPos []int, s string) []int {
	// 方向映射表
	direct := map[byte][2]int{'L': [2]int{0, -1}, 'U': [2]int{-1, 0}, 'R': [2]int{0, 1}, 'D': [2]int{1, 0}}
	// 从起始位置出发上下左右最远可以走的距离
	lmax, umax, rmax, dmax := startPos[1] + 1, startPos[0] + 1, n - startPos[1], n - startPos[0]

	fmt.Printf("lmax: %d, umax: %d, rmax: %d, dmax: %d\n\n", lmax, umax, rmax, dmax)
	pathLen := len(s)

	// 设停止点为(0,0)
	row, col := 0, 0
	rowNext, colNext := map[int]int{0: pathLen}, map[int]int{0: pathLen}
	ret := []int{}

	// 从(0,0)开始从后向前倒推每一步的相对位置，并更新
	for i := pathLen - 1; i >= 0; i-- {
		row, col = row - direct[s[i]][0], col - direct[s[i]][1]
		num := pathLen - i
		fmt.Printf("row: %d, col: %d, num: %d\n", row, col, num)
		fmt.Printf("colNext: %v, rowNext: %v\n", colNext, rowNext)

		/*
		  检查当前位置出发是否会遇上停止点坐标
		  如没有则表示可以走到底(值为pathLen -i)
		  如遭遇则表示最远只能走到那一点(值取该点位置和当前位置之差); 如两个方向同时match,取其中较近的距离
		*/
		if v, ok := colNext[col-lmax]; ok {
			fmt.Printf("col[%d] - lmax[%d] = %d found in colNext!\n", col, lmax, col - lmax)
			num = min(num, v - i - 1)
		}

		if v, ok := rowNext[row-umax]; ok {
			fmt.Printf("row[%d] - umax[%d] = %d found in rowNext!\n", row, umax, row - umax)
			num = min(num, v - i - 1)
		}

		if v, ok := colNext[col+rmax]; ok {
			fmt.Printf("col[%d] + rmax[%d] = %d found in colNext!\n", col, rmax, col + rmax)
			num = min(num, v - i - 1)
		}

		if v, ok := rowNext[row+dmax]; ok {
			fmt.Printf("row[%d] + dmax[%d] = %d found in rowNext!\n", row, dmax, row + dmax)
			num = min(num, v - i - 1)
		}

		fmt.Printf("num after check: %d\n", num)
		// 更新终止点map
		rowNext[row], colNext[col] = i, i
		ret = append([]int{num}, ret...)
	}

	return ret
}

func main() {
	n := []int{3,2,1}
	startPos := [][]int{[]int{0,1}, []int{1,1}, []int{0,0}}
	s := []string{"RRDDLU","LURD","LRUD"}

	for i := range n {
		fmt.Println(executeInstructions(n[i], startPos[i], s[i]))
		fmt.Println("======")
	}
}
