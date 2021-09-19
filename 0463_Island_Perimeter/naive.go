package main

import (
	"fmt"
)

func islandPerimeter(grid [][]int) int {
	// 检查左边和上面的格子
	direct := [][]int{[]int{0,-1}, []int{-1,0}}
	m, n := len(grid), len(grid[0])

	rslt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
		
			rslt += 4
			for k := range direct {
				x, y := i + direct[k][0], j + direct[k][1]
				if x >= 0 && x < m && y >=0 && y < n && grid[x][y] == 1 {
					rslt -= 2
				}
			}
		}
	}

	return rslt
}

func main() {
	grids := [][][]int{[][]int{[]int{0,1,0,0},[]int{1,1,1,0},[]int{0,1,0,0},[]int{1,1,0,0}}, [][]int{[]int{1}}, [][]int{[]int{1,0}}}
	for i := range grids {
		fmt.Println(islandPerimeter(grids[i]))
	}
}
