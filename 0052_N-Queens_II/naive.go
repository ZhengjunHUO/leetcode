package main

import (
	"fmt"
	"strings"
)

var rslt int

func check(curr []string, row, col int) bool {
	for i := row - 1; i >= 0; i-- {
		if curr[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row - 1, col - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if curr[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row - 1, col + 1; i >= 0 && j < len(curr[0]); i, j = i-1, j+1 {
		if curr[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func solve(curr []string, row, n int) {
	if row == n {
		rslt += 1 
		return
	}

	currRow := []byte(strings.Repeat(".", n))

	for col := 0; col < n; col++ {
		if !check(curr, row, col) {
			continue
		}

		currRow[col] = 'Q'
		solve(append(curr, string(currRow)), row+1, n)
		currRow[col] = '.'
	}
}

func totalNQueens(n int) int {
	rslt = 0
	solve([]string{}, 0, n)
	return rslt
}

func main() {
	fmt.Println(totalNQueens(4))
	fmt.Println(totalNQueens(1))
	// 8 queens get 92 solutions
	fmt.Println(totalNQueens(8))
}
