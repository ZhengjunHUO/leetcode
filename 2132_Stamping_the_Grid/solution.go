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

func calcPrefixSum(grid *[][]int, sum *[][]int) {
	m, n := len(*grid), len((*grid)[0])
	*sum = make([][]int, m+1)
	for i := range *sum {
		(*sum)[i] = make([]int, n+1)
	}

	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			(*sum)[i][j] = (*sum)[i][j-1] + (*sum)[i-1][j] - (*sum)[i-1][j-1] + (*grid)[i-1][j-1]
		}
	}
}

func rangeSum(grid *[][]int, row1, col1, row2, col2 int) int {
	return (*grid)[row2+1][col2+1] - (*grid)[row2+1][col1] - (*grid)[row1][col2+1] + (*grid)[row1][col1]
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m, n := len(grid), len(grid[0])

	var sum [][]int
	calcPrefixSum(&grid, &sum)
	//fmt.Println("sum:")
	//fmt.Println(sum)

	mark := make([][]int, m)
	for i := range mark {
		mark[i] = make([]int, n)
	}

	for i := stampHeight-1; i < m; i++ {
		for j := stampWidth-1; j < n; j++ {
			//fmt.Printf("rangeSum &sum, %d, %d, %d, %d\n", i-stampHeight+1, j-stampWidth+1, i, j)
			if rangeSum(&sum, i-stampHeight+1, j-stampWidth+1, i, j) == 0 {
				mark[i][j] = 1
			}
		}
	}
	//fmt.Println("mark:")
	//fmt.Println(mark)

	var markSum [][]int
	calcPrefixSum(&mark, &markSum)
	//fmt.Println("markSum:")
	//fmt.Println(markSum)

	for i := range grid {
		for j := range grid[i] {
			//fmt.Printf("rangeSum &markSum, %d, %d, %d, %d\n", i, j, min(i+stampHeight-1,m-1), min(j+stampWidth-1,n-1))
			if grid[i][j] == 0 && rangeSum(&markSum, i, j, min(i+stampHeight-1,m-1), min(j+stampWidth-1,n-1)) == 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	grids := [][][]int{[][]int{[]int{1,0,0,0},[]int{1,0,0,0},[]int{1,0,0,0},[]int{1,0,0,0},[]int{1,0,0,0}}, [][]int{[]int{1,0,0,0},[]int{0,1,0,0},[]int{0,0,1,0},[]int{0,0,0,1}}}
	h := []int{4,2}
	w := []int{3,2}
	for i := range grids {
		fmt.Println(possibleToStamp(grids[i], h[i], w[i]))
	}
}
