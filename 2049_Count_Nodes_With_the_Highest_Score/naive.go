package main

import (
	"fmt"
)

func updateMax(a int, max, count *int) {
	if a > (*max) {
		*max, *count = a, 1
	}else if a == (*max) {
		*count = (*count)+1
	}
}

func calculateSize(curr int, children map[int][]int, size *[]int) int {
	if (*size)[curr] != 0 {
		return (*size)[curr]
	}

	var ret int
	if v, ok := children[curr]; !ok {
		(*size)[curr] = 1
		return 1
	}else{
		ret = 1
		for j := range v {
			ret += calculateSize(v[j], children, size)
		}
	}

	(*size)[curr] = ret
	return ret
}

func countHighestScoreNodes(parents []int) int {
	children := make(map[int][]int)
	size := make([]int, len(parents))
	maxscore, count := 0, 0

	for i := range parents {
		if parents[i] != -1 {
			if _, ok := children[parents[i]]; ok {
				children[parents[i]] = append(children[parents[i]], i)
			}else{
				children[parents[i]] = []int{i}
			}
		}
	}

	for i := range parents {
		if size[i] != 0 {
			continue
		}

		calculateSize(i, children, &size)
	}

	for i := range parents {
		if v, ok := children[i]; !ok {
			// leaf
			updateMax(len(parents)-1, &maxscore, &count)
		}else{
			val := 1
			for j := range v {
				val *= size[v[j]]
			}

			if parents[i] == -1 {
				// root
				updateMax(val, &maxscore, &count)
			}else{
				// intermediate
				val *= (size[parents[i]] - size[i])
				updateMax(val, &maxscore, &count)
			}
		}
	}

	return count
}

func main() {
	parents := [][]int{[]int{-1,2,0,2,0},[]int{-1,2,0}}
	for i := range parents {
		fmt.Println(countHighestScoreNodes(parents[i]))
	}
}
