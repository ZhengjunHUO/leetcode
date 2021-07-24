package main

import (
	"fmt"
)

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	prim := []int{2, 3, 5}
	for _, v := range prim {
		for n%v == 0 {
			n /= v
		}
	}

	return n == 1    
}

func main() {
	nums := []int{6, 8, 14, 1}
	for _, v := range nums {
		fmt.Println(isUgly(v))
	}
}
