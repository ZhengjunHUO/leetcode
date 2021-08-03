package main

import (
	"fmt"
)

func backtrack(grid [][]int, min *int, x, y, curr int) {
	if x == len(grid) - 1 && y == len(grid[0]) - 1 {
		if curr < *min {
			*min = curr
		}
	}

	if x < len(grid) - 1 {
		backtrack(grid, min, x+1, y, curr + grid[x+1][y])
	}
	if y < len(grid[0]) - 1 {
		backtrack(grid, min, x, y+1, curr + grid[x][y+1])
	}
}

func minPathSum(grid [][]int) int {
	min := 40000
	backtrack(grid, &min, 0, 0, grid[0][0])
    	return min
}

func main() {
	grids := [][][]int{[][]int{[]int{1,3,1},[]int{1,5,1},[]int{4,2,1}}, [][]int{[]int{1,2,3},[]int{4,5,6}}}
	for _,v := range grids {
		fmt.Println(minPathSum(v))
	}
}
