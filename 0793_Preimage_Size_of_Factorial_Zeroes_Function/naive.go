package main

import (
	"fmt"
)

// trailingZeroes(MaxUint32) => 1073741816 > 10^9 可以作为右边界
const MaxUint32 = 1<<32 - 1

// 参考0172_Factorial_Trailing_Zeroes
func trailingZeroes(n int) int {
        if n == 0 {
                return 0
        }

        return n/5 + trailingZeroes(n/5)
}

func preimageSizeFZF(k int) int {
	// 求x取值的左边界
	l, r := 0, MaxUint32
	for l < r {
		m := l + (r-l)/2

		if trailingZeroes(m) >= k {
			r = m
		}else{
			l = m+1
		}
	}

	leftx := l

	// 求x取值的右边界
	l, r = 0, MaxUint32
	for l < r {
		m := l + (r-l)/2

		if trailingZeroes(m) <= k {
			l = m+1
		}else{
			r = m
		}
	}

	rightx := l - 1

	fmt.Println(leftx, rightx)

	return rightx - leftx + 1
}

func main() {
	ks := []int{0,5,3}
	for i := range ks {
		fmt.Println(preimageSizeFZF(ks[i]))
	}

	fmt.Println(preimageSizeFZF(23))
}
