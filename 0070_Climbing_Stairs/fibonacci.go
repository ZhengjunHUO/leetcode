package main

import (
	"fmt"
)

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	f1, f2, rslt := 1, 2, 0
	for i:=3; i<=n; i++ {
		rslt = f1 + f2
		f1, f2 = f2, rslt
	}

	return rslt
}

func main() {
	ns := []int{2,3,5,8}
	for _,v := range ns {
		fmt.Println(climbStairs(v))
	}
}
