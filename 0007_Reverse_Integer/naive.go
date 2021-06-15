package main

import (
	"fmt"
)

// (2^31 - 1) / 10 = 2147483647 / 10 
const INTMAX int = 214748364
// (- 2 ^ 31) / 10 = -2147483648 / 10
const INTMIN int = -214748364

// 时间复杂度是O(logN)
func reverse(x int) int {
	var curr, rslt int 
	for x != 0 {
		curr = x%10

		if (rslt > INTMAX) || (rslt == INTMAX && curr > 7) || (rslt < INTMIN) || (rslt == INTMIN && curr < -8) {
			return 0
		}
		rslt = rslt*10 + curr
		x /= 10
	} 

	return rslt
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))

	// reverse of (2^31 - 1)
	fmt.Println(reverse(7463847412))
	// 超出最大正数范围，返回0
	fmt.Println(reverse(8463847412))
}
