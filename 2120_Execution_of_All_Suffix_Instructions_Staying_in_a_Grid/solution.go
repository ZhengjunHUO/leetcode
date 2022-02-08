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

func executeInstructions(n int, startPos []int, s string) []int {
	direct := map[byte][2]int{'L': [2]int{0, -1}, 'U': [2]int{-1, 0}, 'R': [2]int{0, 1}, 'D': [2]int{1, 0}}
	lmax, umax, rmax, dmax := startPos[1] + 1, startPos[0] + 1, n - startPos[1], n - startPos[0]
	pathLen := len(s)

	row, col := 0, 0
	rowNext, colNext := map[int]int{0: pathLen}, map[int]int{0: pathLen}
	ret := []int{}

	for i := pathLen - 1; i >= 0; i-- {
		row, col = row - direct[s[i]][0], col - direct[s[i]][1]
		num := pathLen - i

		if v, ok := colNext[col-lmax]; ok {
			num = min(num, v - i - 1)
		}

		if v, ok := rowNext[row-umax]; ok {
			num = min(num, v - i - 1)
		}

		if v, ok := colNext[col+rmax]; ok {
			num = min(num, v - i - 1)
		}

		if v, ok := rowNext[row+dmax]; ok {
			num = min(num, v - i - 1)
		}

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
	}
}
