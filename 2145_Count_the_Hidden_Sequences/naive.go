package main

import (
	"fmt"
)

func updateMin(num int, min *int) {
	if num < *min {
		*min = num
	}
}

func updateMax(num int, max *int) {
	if num > *max {
		*max = num
	}
}

func numberOfArrays(differences []int, lower int, upper int) int {
	curr, min, max := 0, 10001, -10001
	
	// 假设第一个元素为0，根据差值数列计算出原数列的最值，可以得到数列的实际波动范围
	for i := range differences {
		next := curr + differences[i]
		if next > curr {
			updateMin(curr, &min)
			updateMax(next, &max)
		}else{
			updateMin(next, &min)
			updateMax(curr, &max)
		}
		curr = next
	}

	// lower和upper决定了数列波动的最大范围
	// 用最大范围减去实际的波动范围就是可能的原数列数量
	ret := (upper - lower + 1) - (max - min)

	// 如果实际波动范围大于给出的最大范围则返回0
	if ret < 0 {
		return 0
	}

	return ret
}

func main() {
	diffs := [][]int{[]int{1,-3,4}, []int{3,-4,5,1,-2}, []int{4,-7,2}}
	lowers := []int{1, -4, 3}
	uppers := []int{6, 5, 6}

	for i := range diffs {
		fmt.Println(numberOfArrays(diffs[i], lowers[i], uppers[i]))
	}
}
