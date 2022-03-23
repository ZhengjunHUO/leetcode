package main

import (
	"fmt"
)

const MAX_INT64 = int64(^uint64(0) >> 1)

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

// solution from @votrubac
func gridGame(grid [][]int) int64 {
	ret := MAX_INT64
	var above, bellow int64

	for _,v := range grid[0] {
		above += int64(v)
	}

	for i := range grid[0] {
		above -= int64(grid[0][i])
		ret = min(ret, max(above, bellow))
		bellow += int64(grid[1][i])
	}

	return ret
}

func main() {
	grid := [][][]int{[][]int{[]int{2,5,4},[]int{1,5,1}}, [][]int{[]int{3,3,1},[]int{8,5,2}}, [][]int{[]int{1,3,1,15},[]int{1,3,3,1}}}
	for i := range grid {
		fmt.Println(gridGame(grid[i]))
	}
}
