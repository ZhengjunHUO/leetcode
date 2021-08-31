package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	return ((n - 1) & n ) == 0
}

func main() {
	ns := []int{1,16,3,4,5}
	for i := range ns {
		fmt.Println(isPowerOfTwo(ns[i]))
	}
}
