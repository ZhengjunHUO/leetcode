package main

import (
	"fmt"
)

var dict [][]int

func count(l, r int) int {
	// nil也算是一种情况；只有一个结点也只有一种情况
	if l >= r {
		return 1
	}

	// 查看cache
	if dict[l][r] != 0 {
		//fmt.Println("[INFO] Hit cache!")
		return dict[l][r]
	}

	rslt := 0
	// 遍历区间内的结点，计算分别以每个结点为根能有几种组合，累加至rslt
	for i:=l; i<=r; i++ {
		rslt += count(l, i-1) * count(i+1, r)
	}
	// 升级cache
	dict[l][r] = rslt

	return rslt
}

func numTrees(n int) int {
	dict = make([][]int, n+1)
	for i := range dict {
		dict[i] = make([]int, n+1)
	}

	return count(1,n)
}

func main() {
	ns := []int{3,1,6,7}
	for _, v := range ns {
		fmt.Println(numTrees(v))
	}
}
