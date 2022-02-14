package main

import (
	"fmt"
)

type NumMatrix struct {
	m, n int
	table [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	// padding an extra zeroed column and row
	table := make([][]int, m+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}

	// calculate accumulated sum matrix, table[i][j] is the sum between matrix[0][0] and matrix[i-1][j-1]
	// "- table[i-1][j-1]" because "table[i-1][j] + table[i][j-1]" add 2 times table[i-1][j-1]
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			table[i][j] = table[i-1][j] + table[i][j-1] - table[i-1][j-1] + matrix[i-1][j-1]
			//fmt.Printf("table[%d][%d] = %d + %d - %d + %d = %d\n", i, j, table[i-1][j], table[i][j-1], table[i-1][j-1], matrix[i-1][j-1], table[i][j])
		}
	}

	//fmt.Println(table)

	numMatrix := NumMatrix {
		m: m,
		n: n,
		table: table,
	}

	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// this.table got an extra padding, pay attention to the index
	// "+ this.table[row1][col1]" because "- this.table[row1][col2+1] - this.table[row2+1][col1]" minus 2 times this.table[row1][col1]
	return this.table[row2+1][col2+1] - this.table[row1][col2+1] - this.table[row2+1][col1] + this.table[row1][col1]
}

func main() {
	numMatrix := Constructor([][]int{[]int{3, 0, 1, 4, 2}, []int{5, 6, 3, 2, 1}, []int{1, 2, 0, 1, 5}, []int{4, 1, 0, 1, 7}, []int{1, 0, 3, 0, 5}})
	fmt.Println(numMatrix.SumRegion(2, 1, 4, 3))
	fmt.Println(numMatrix.SumRegion(1, 1, 2, 2))
	fmt.Println(numMatrix.SumRegion(1, 2, 2, 4))
}
