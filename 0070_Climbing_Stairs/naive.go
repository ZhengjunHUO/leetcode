package main

import (
	"fmt"
)

func backtrack(n, curr int, cnt *int) {
	if curr == n {
		*cnt += 1
		return
	}

	if curr < n {
		backtrack(n, curr+1, cnt)
	}
	if curr < n-1 {
		backtrack(n, curr+2, cnt)
	}
}

func climbStairs(n int) int {
	cnt := 0
	backtrack(n, 0, &cnt)
	return cnt    
}

func main() {
	ns := []int{2,3,5,8}
	for _,v := range ns {
		fmt.Println(climbStairs(v))
	}
}
