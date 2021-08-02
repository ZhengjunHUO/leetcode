package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	rslt := make([][]int, n)
	for i := range rslt {
		rslt[i] = make([]int, n)
	}

	movMatrix := [][]int{[]int{0,1}, []int{1,0}, []int{0,-1}, []int{-1, 0}}
	movIdx := 0
	steps, row, col, val := []int{n, n-1}, 0, -1, 1

	for steps[movIdx%2]!=0 {
		for i:=0; i<steps[movIdx%2]; i++ {
			row, col = row+movMatrix[movIdx][0], col+movMatrix[movIdx][1]
			rslt[row][col] = val
			val++
		}

		steps[movIdx%2]--
		movIdx = (movIdx + 1)%4
	}	

	return rslt
}

func main() {
	ns := []int{3,1}
	for _,v := range ns {
		fmt.Println(generateMatrix(v))
	}
}
