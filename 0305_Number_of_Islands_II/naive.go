package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func numIslands2(m,n int, positions [][]int) []int {
	uf := graph.NewUF(m*n)
	uf.SetCount(0)

	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	rslt := []int{}

	dirs := [][]int{[]int{0,-1}, []int{1,0}, []int{0,1}, []int{-1,0}}

	for i := range positions {
		matrix[positions[i][0]][positions[i][1]] = 1
		uf.SetCount(uf.Count() + 1)
		for k := 0; k < 4; k++ {
			x, y := positions[i][0]+dirs[k][0], positions[i][1]+dirs[k][1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if matrix[x][y] == 1 {
					uf.Union(positions[i][0]*n+positions[i][1], x*n+y)
				}
			}
		}

		rslt = append(rslt, uf.Count())
	}

	return rslt
}

func main() {
	fmt.Println(numIslands2(3, 3, [][]int{[]int{0,0}, []int{0,1}, []int{1,2}, []int{2,1}, []int{1,1}}))
}
