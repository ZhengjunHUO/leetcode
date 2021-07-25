package main

import (
	"fmt"
	"math"
)

func backtrack(dict []int, rslt *int, rest, index, count int) {
	if rest == 0 {
		if count < *rslt {
			*rslt = count
		}
		return 
	}

	if index >= len(dict) {
		return
	}

	for i, j := rest, 0; i>=0; i, j = i - dict[index], j+1 {
		backtrack(dict, rslt, i, index+1, count+j)
	}
}

func numSquares(n int) int {
	cnt := n

	temp := int(math.Sqrt(float64(n)))
	dict := make([]int, temp)
	for i:=temp; i>0; i-- {
		dict[temp-i] = i*i
	}
	
	backtrack(dict, &cnt, n, 0, 0)
	return cnt
}

func main() {
	nums := []int{12, 13, 11, 10}
	for _, v := range nums {
		fmt.Println(numSquares(v))
	}
}
