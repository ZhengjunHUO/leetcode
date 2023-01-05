package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func largestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	uf := graph.NewUF(m*n)

	// 构建初始岛屿分布图
	dirs := [][]int{[]int{0,-1}, []int{1,0}, []int{0,1}, []int{-1,0}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				for k := 0; k < 4; k++ {
					x, y := i+dirs[k][0], j+dirs[k][1]
					if x >= 0 && x < m && y >= 0 && y < n {
						if grid[x][y] == 1 {
							uf.Union(i*n+j, x*n+y)
						}
					}
				}
			}
		}
	}

	// 保存UF
	c, p, s := uf.Count(), uf.Parent(), uf.Size()
	largest := getMax(s)
	//fmt.Println(uf.Parent(), uf.Size())

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 可能的填海地点
			if grid[i][j] == 0 {
				// Restore初始地图
				uf.SetCount(c)
				uf.SetParent(p)
				uf.SetSize(s)

				for k := 0; k < 4; k++ {
					x, y := i+dirs[k][0], j+dirs[k][1]
					if x >= 0 && x < m && y >= 0 && y < n {
						if grid[x][y] == 1 {
							uf.Union(i*n+j, x*n+y)
						}
					}
				}

				// 观察填海后是否造出更大的岛屿
				if temp := getMax(uf.Size()); temp > largest {
					largest = temp
				}
			}
		}
	}

	//fmt.Println(uf.Parent(), uf.Size())
	return largest
}

func getMax(size []int) int {
	max := 0
	for i := range size {
		if size[i] > max {
			max = size[i]
		}
	}

	return max
}

func main() {
	grids := [][][]int{[][]int{[]int{1, 0}, []int{0, 1}}, [][]int{[]int{1, 1}, []int{1, 0}}, [][]int{[]int{1, 1}, []int{1, 1}}, [][]int{[]int{1}}, [][]int{[]int{0}}, [][]int{[]int{1,1,0,0,0}, []int{1,1,0,0,0}, []int{0,0,1,0,0}, []int{0,0,0,1,1}}}
	for i := range grids {
		fmt.Println(largestIsland(grids[i]))
	}
}
