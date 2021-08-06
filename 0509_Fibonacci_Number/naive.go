package main

import (
	"fmt"
)

func fib(n int) int {
	if n < 2 {
		return n
	} 

	curr, preprev, prev := 0, 0, 1
	for i := 2; i <= n; i++ {
		curr = prev + preprev
		preprev, prev = prev, curr
	}

	return curr
}

func main() {
	ns := []int{2,3,4,10}
	for _,v := range ns {
		fmt.Println(fib(v))
	}
}
