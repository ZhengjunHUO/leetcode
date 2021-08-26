package main

import (
	"fmt"
)

func recursive(tab [][]int, min *int, currPrice, currPos, dst, k int) {
	if k >= 0 && currPos == dst {
		if currPrice < *min {
			*min = currPrice
		}
		return
	}

	for k == 0 {
		return 
	}

	for j := 0; j < len(tab); j++ {
		if tab[currPos][j] != 0 {
			recursive(tab, min, currPrice+tab[currPos][j], j, dst, k-1)
		}
	}
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	tab := make([][]int, n)
	for i := range tab {
		tab[i] = make([]int, n)
	}
	for i := range flights {
		tab[flights[i][0]][flights[i][1]] = flights[i][2]
	}

	min := 1000001	
	recursive(tab, &min, 0, src, dst, k+1)
	
	if min == 1000001 {
		return -1
	}

	return min
}

func main() {
	f := [][]int{[]int{0,1,100},[]int{1,2,100},[]int{0,2,500}}
	fmt.Println(findCheapestPrice(3, f, 0, 2, 1))
	fmt.Println(findCheapestPrice(3, f, 0, 2, 0))
}
