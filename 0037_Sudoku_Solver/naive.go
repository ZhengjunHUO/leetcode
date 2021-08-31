package main

import (
	"fmt"
)

var found bool

func check(board [][]byte, row, col int, num byte) bool {
	n := 3
	rowBlock, colBlock := row/n, col/n

	// 检查九宫格
	for i:=n*rowBlock; i<n*rowBlock+n; i++ {
		for j:=n*colBlock; j<n*colBlock+n; j++ {
			if board[i][j] == num {
				if !(i==row && j==col) {
					return false
				}
			}
		}
	}

	// 检查横向
	for j:=0; j<len(board[0]); j++ {
		if board[row][j] == num {
			if j != col {
				return false
			} 
		}
	}

	// 检查纵向
	for i:=0; i<len(board); i++ {
		if board[i][col] == num {
			if i != row {
				return false
			}
		}
	}

	return true
}

func solve(board [][]byte, boxNum int) {
	n := 9
	maxNum := 8*n + 8 
	
	// 找到下一个空格
	for boxNum <= maxNum && board[boxNum/n][boxNum%n] != '.' {
		boxNum++
	}

	if boxNum > maxNum {
		found = true
		return
	}	

	row, col := boxNum/n, boxNum%n
	
	// 尝试填入数字
	for i:=1; i<=9; i++ {
		if !check(board, row, col, byte(48+i)) {
			continue
		} 
		
		board[row][col] = byte(48+i)
		solve(board, boxNum)
		if found == true {
			break
		}
		board[row][col] = '.'
	}

}

func solveSudoku(board [][]byte)  {
	found = false
	solve(board, 0)
}

func main() {
	board := [][]byte{[]byte{'5','3','.','.','7','.','.','.','.'},[]byte{'6','.','.','1','9','5','.','.','.'},[]byte{'.','9','8','.','.','.','.','6','.'},[]byte{'8','.','.','.','6','.','.','.','3'},[]byte{'4','.','.','8','.','3','.','.','1'},[]byte{'7','.','.','.','2','.','.','.','6'},[]byte{'.','6','.','.','.','.','2','8','.'},[]byte{'.','.','.','4','1','9','.','.','5'},[]byte{'.','.','.','.','8','.','.','7','9'}}

	fmt.Println("Problem: ")
	for i := range board {
		fmt.Printf("%s\n", board[i])
	}

	solveSudoku(board)

	fmt.Println("\nSolution: ")
	for i := range board {
		fmt.Printf("%s\n", board[i])
	}
}
