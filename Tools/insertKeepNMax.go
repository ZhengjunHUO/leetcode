package main

import (
	"fmt"
	"sort"
)

func main() {
	// case 0
	//x, y := []int{2,3,4,7,10}, 1

	// case 1
	//x, y := []int{2,4,5,7,10}, 3

	x, y := []int{0,0,0,0}, 6
	//x, y := []int{2,3,4,7,10}, 6
	
	idx := sort.SearchInts(x, y)
	fmt.Println(idx)
	
	switch idx {
	case 0:
		fmt.Println("pass")
	case 1:
		x[0] = y
		fmt.Println("the smallest")
		fmt.Println(x)
	default:
		fmt.Println("shift array")
		copy(x[0:idx-1], x[1:idx])
		x[idx-1] = y
		fmt.Println(x)
	}
}
