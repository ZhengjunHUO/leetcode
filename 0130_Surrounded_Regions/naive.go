package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	uf := godtype.NewUF(m*n+1)
	directions := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}

	notFlipRoot := m*n
	for i := range board {
		if board[i][0] == 'O' {
			uf.Union(notFlipRoot, i*n)
		}
		if board[i][n-1] == 'O' {
			uf.Union(notFlipRoot, i*n + n-1)
		}
	}	
	
	for i := range board[0] {
		if board[0][i] == 'O' {
			uf.Union(notFlipRoot, i)
		}
		if board[m-1][i] == 'O' {
			uf.Union(notFlipRoot, (m-1)*n+i)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x, y := i + directions[k][0], j + directions[k][1]
					if board[x][y] == 'O' {
						uf.Union(i*n+j, x*n+y)
					} 
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if ! uf.IsLinked(notFlipRoot, i*n+j) {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{[]byte{'X','X','X','X'},[]byte{'X','O','O','X'},[]byte{'X','X','O','X'},[]byte{'X','O','X','X'}}
	fmt.Printf("Before: %s\n", board)
	solve(board)
	fmt.Printf("After:  %s\n", board)
}
