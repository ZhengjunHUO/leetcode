package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	uf := godtype.NewUF(m*n)
	directions := [][]int{[]int{0,0}, []int{1,0}, []int{1,1}, []int{0,1}}
	
	// 使用一个2*2的框来遍历矩阵
	for i := 0; i < m - 1; i++ {
		for j := 0; j < n - 1; j++ {
			// 以左下右上的顺序遍历框中的四个元素，如果当前元素和下一个元素都是1则调用Union方法
			for k := 0; k < 4; k ++ {
				if grid[i+directions[k][0]][j+directions[k][1]] == 1 {
					if grid[i+directions[(k+1)%4][0]][j+directions[(k+1)%4][1]] == 1 {
						uf.Union((i+directions[k][0])*n+j+directions[k][1], (i+directions[(k+1)%4][0])*n+j+directions[(k+1)%4][1])
					}
				}
			}
		}
	}

	dict := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// 同一个岛的1指向同一个根结点
				dict[uf.FindRoot(i*n+j)] += 1
			}	
		}
	}
	
	//fmt.Println(dict)
	max := 0
	for _, v := range dict {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	grid := [][]int{[]int{0,0,1,0,0,0,0,1,0,0,0,0,0},[]int{0,0,0,0,0,0,0,1,1,1,0,0,0},[]int{0,1,1,0,1,0,0,0,0,0,0,0,0},[]int{0,1,0,0,1,1,0,0,1,0,1,0,0},[]int{0,1,0,0,1,1,0,0,1,1,1,0,0},[]int{0,0,0,0,0,0,0,0,0,0,1,0,0},[]int{0,0,0,0,0,0,0,1,1,1,0,0,0},[]int{0,0,0,0,0,0,0,1,1,0,0,0,0}}
	fmt.Println(maxAreaOfIsland(grid))
}
