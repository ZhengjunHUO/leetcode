package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	return n/5 + trailingZeroes(n/5)
}

func main() {
	ns := []int{3,5,0,20,25}
	// 20: 4个0; 25: 6个0
	for i := range ns {
		fmt.Println(trailingZeroes(ns[i]))
	}
	
	x := 1<<32 - 1
	fmt.Println(trailingZeroes(x))
}
