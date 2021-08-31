package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	count := 0
	for n > 0 {
		count += (n & 1)
		n = n >> 1
	}

	return count == 1
}

func main() {
	ns := []int{1,16,3,4,5}
	for i := range ns {
		fmt.Println(isPowerOfTwo(ns[i]))
	}
}
