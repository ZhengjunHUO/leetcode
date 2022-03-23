package main

import (
	"fmt"
)

func chooseRoute(grid [][]int) (int, int) {
	currVal := grid[0][0]
	for _,v := range grid[1] {
		currVal += v
	}

	currIdx, maxIdx, maxVal := 1, 0, currVal
	for currIdx < len(grid[1]) {
		currVal += grid[0][currIdx]
		currVal -= grid[1][currIdx-1]
		if currVal > maxVal {
			maxIdx, maxVal = currIdx, currVal
		}
		currIdx++
	}

	return maxIdx, maxVal
}

// naive implementation, simulate robots' movement
func gridGame(grid [][]int) int64 {
	// for robot 1, pick up an index to go below with maximum points
	maxIdx, _ := chooseRoute(grid)

	// put zero to the choosen path
	for i:=0; i<=maxIdx; i++ {
		grid[0][i] = 0
	}
	for i:=maxIdx; i<len(grid[1]); i++ {
		grid[1][i] = 0
	}

	// for robot 2 do the same thing
	_, ret := chooseRoute(grid)
	return int64(ret)
}

func main() {
	grid := [][][]int{[][]int{[]int{2,5,4},[]int{1,5,1}}, [][]int{[]int{3,3,1},[]int{8,5,2}}, [][]int{[]int{1,3,1,15},[]int{1,3,3,1}}}
	for i := range grid {
		fmt.Println(gridGame(grid[i]))
	}
}
