package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	m, n, count := len(s)-1, len(g)-1, 0

	for m >= 0 && n >= 0 {
		if s[m] >= g[n] {
			count++
			m, n = m - 1, n - 1
		}else{
			n -= 1
		}
	}

	return count
}

func main() {
	g := [][]int{[]int{1,2,3}, []int{1,2}}
	s := [][]int{[]int{1,1}, []int{1,2,3}}

	for i := range g {
		fmt.Println(findContentChildren(g[i], s[i]))
	}
}
