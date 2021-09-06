package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	cumsum	[]int
	max	int
}

/*
  按照权重数列构建一个累加数列
  按随机值坐落的区间来决定返回哪个index
*/
func Constructor(w []int) Solution {
	cumsum := make([]int, len(w))
	cumsum[0] = w[0]
	
	for i := 1; i < len(w); i++ {
		cumsum[i] = cumsum[i-1] + w[i]
	}

	return Solution{cumsum, cumsum[len(cumsum)-1]}
}

func (this *Solution) PickIndex() int {
	p := rand.Intn(this.max) + 1

	l, r := 0, len(this.cumsum)
	for l < r {
		m := l + (r-l)/2
		if this.cumsum[m] >= p {
			r = m
		}else{
			l = m + 1	
		}
	}

	return l
}

func main() {
	obj := Constructor([]int{4,2,8,6})
	dict := make(map[int]int)

	for i := 0; i < 20000; i++ {
		dict[obj.PickIndex()]++
	}

	fmt.Println(dict)
}
