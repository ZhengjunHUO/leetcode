package main

import (
	"fmt"
)

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	var ret,diffRow,diffCol int
	startRow, startCol, endRow, endCol := startPos[0], startPos[1], homePos[0], homePos[1] 
	if startRow < endRow {
		diffRow = 1
	}else if startRow > endRow {
		diffRow = -1
	}

	if startCol < endCol {
		diffCol = 1
	}else if startCol > endCol {
		diffCol = -1
	}

	for startRow != endRow {
		startRow += diffRow
		ret += rowCosts[startRow]
	}

	for startCol != endCol {
		startCol += diffCol
		ret += colCosts[startCol]
	}

	return ret
}

func main() {
	startPos := [][]int{[]int{1, 0}, []int{0, 0}}
	homePos := [][]int{[]int{2, 3}, []int{0, 0}}
	rowCosts := [][]int{[]int{5, 4, 3}, []int{5}}
	colCosts := [][]int{[]int{8, 2, 6, 7}, []int{26}}

	for i := range startPos {
		fmt.Println(minCost(startPos[i], homePos[i], rowCosts[i], colCosts[i]))
	}
}
