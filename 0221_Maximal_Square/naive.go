package main

import (
	"fmt"
)

/*
  从右到左，从下到上扫描matrix，记录以matrix[i][j]为左上角的正方形的最大面积
*/
func maximalSquare(matrix [][]byte) int {
	max, row, col := 0, len(matrix), len(matrix[0])

	table := make([]int, col)
	for i := range table {
		if matrix[row-1][i] == '0' {
			table[i] = 0
		}else{
			table[i] = 1
		}
	}
	//fmt.Println(table)
	
	for i:=row-2; i>=0; i-- {
		for j:=col-1; j>=0; j-- {
			if matrix[i][j] == '0' {
				table[j] = 0
				continue
			}

			if table[j] == 0 || j == col - 1 {
				table[j] = int(matrix[i][j]-'0')
				// update max
				if table[j] > max {
					max = table[j]
				}
				continue
			}
	
			/*
			  此时matrix[i][j] == '1', table[j] > 0
			  需要检查[i][j]->[i][j+x], [i][j+x]->[i+x][j+x], 需要注意边界
			*/
			if table[j] + j > col - 1 {
				table[j] = col - 1
				continue
			}

			canExtend := true
			for k := 1; k <= table[j]; k++ {
				if matrix[i][j+k] == '0' || matrix[i+k][j+table[j]] == '0' {
					canExtend = false
					break		
				}
			}

			if canExtend == true {
				table[j] += 1
				if table[j] > max {
					max = table[j]
				}
			}else{
				table[j] = 1
			}
		}
		//fmt.Println(table)
	}

	return max*max
}

func main() {
	matrix := [][][]byte{[][]byte{[]byte{'1','0','1','0','0'},[]byte{'1','0','1','1','1'},[]byte{'1','1','1','1','1'},[]byte{'1','0','0','1','0'}}, [][]byte{[]byte{'0','1'},[]byte{'1','0'}}, [][]byte{[]byte{'0'}}}
	for i := range matrix {
		fmt.Println(maximalSquare(matrix[i]))
	}
}
