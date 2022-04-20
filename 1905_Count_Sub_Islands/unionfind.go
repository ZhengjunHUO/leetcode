package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

// connect all the cells
func buildUF(grid [][]int, uf *godtype.UF) {
	m, n := len(grid), len(grid[0])

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if grid[i][j] == 1 {
				// look at the right cell
				if i+1<m && grid[i+1][j] == 1 {
					uf.Union(i*m+j, (i+1)*m+j)
				}

				// look at the down cell
				if j+1<n && grid[i][j+1] == 1 {
					uf.Union(i*m+j, i*m+j+1)
				}
			}else{
				// if cell is "water", put -1 as parent
				parent := uf.Parent()
				parent[i*m+j] = -1
				uf.SetParent(parent)
			}
		}
	}

	// make sure the parent of all "land" cells is the root parent
	parent := uf.Parent()
	for i := range parent {
		if parent[i] != -1 {
			root := uf.FindRoot(parent[i])
			parent[i] = root
		}
	}
	uf.SetParent(parent)
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	uf1, uf2 := godtype.NewUF(m*n), godtype.NewUF(m*n)

	buildUF(grid1, uf1)
	buildUF(grid2, uf2)

	// find out all islands in grid2 based on result of UnionFind
	island2 := make(map[int][]int)
	for i, v := range uf2.Parent() {
		if v != -1 {
			// group all the "land" cells 
			if _, ok := island2[v]; ok {
				island2[v] = append(island2[v], i)
			}else{
				island2[v] = []int{i}
			}
		}
	}

	var ret int
	parent1 := uf1.Parent()
	// v is []int representing all the cells in a sub island (grid2)
	for _, v := range island2 {
		island1 := make(map[int]struct{})
		// for each cell of the sub island (grid2), check the cell in the same position in the grid1
		for _, val := range v {
			// if the corresponding cell in grid1 is "water", skip this sub island
			if parent1[val] == -1 {
				goto checkNextIsland
			}
			// if the corresponding cell in grid1 is "land", find its "root" in grid and add to a set
			island1[parent1[val]] = struct{}{}
		}

		// if all corresponding cells in grid1 are "land", 1 sub island is found
		if len(island1) == 1 {
			ret++
		}
		checkNextIsland:
	}

	return ret
}

func main() {
	grid1 := [][][]int{[][]int{[]int{1,1,1,0,0},[]int{0,1,1,1,1},[]int{0,0,0,0,0},[]int{1,0,0,0,0},[]int{1,1,0,1,1}}, [][]int{[]int{1,0,1,0,1},[]int{1,1,1,1,1},[]int{0,0,0,0,0},[]int{1,1,1,1,1},[]int{1,0,1,0,1}}}
	grid2 := [][][]int{[][]int{[]int{1,1,1,0,0},[]int{0,0,1,1,1},[]int{0,1,0,0,0},[]int{1,0,1,1,0},[]int{0,1,0,1,0}}, [][]int{[]int{0,0,0,0,0},[]int{1,1,1,1,1},[]int{0,1,0,1,0},[]int{0,1,0,1,0},[]int{1,0,0,0,1}}}

	for i := range grid1 {
		fmt.Println(countSubIslands(grid1[i], grid2[i]))
	}
}
