package main

import (
	"fmt"
)

func findFarmland(land [][]int) [][]int {
	ret := [][]int{}
	rows, cols := len(land), len(land[0])

	for r := range land {
		for c := range land[r] {
			// check only the farmland
			if land[r][c] == 1 {
				i, j := r, c
				// since all the farmlands are rectangular areas and
				// there's no adjacent group, we can simply check the right and bottom boundary
				for i<rows && land[i][c] == 1 {
					i++
				}
				for j<cols && land[r][j] == 1 {
					j++
				}

				// update the result
				ret = append(ret, []int{r,c,i-1,j-1})
				// zero the group which is already taken into consideration
				for x:=r; x<i; x++ {
					for y:=c; y<j; y++ {
						land[x][y] = 0
					}
				}
			}
		}
	}

	return ret
}

func main() {
	land := [][][]int{[][]int{[]int{1,0,0},[]int{0,1,1},[]int{0,1,1}}, [][]int{[]int{1,1},[]int{1,1}}, [][]int{[]int{0}}}
	for i := range land {
		fmt.Println(findFarmland(land[i]))
	}
}
