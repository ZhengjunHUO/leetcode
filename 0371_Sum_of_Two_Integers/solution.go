package main

import (
	"fmt"
)

/*
@kishynivas10给出了很好的解释，观察XOR可以发现不考虑进位的情况下，其得到的值是符合加法规律的
对于二进制只有两个数都是1的时候才需要进位，而与操作(&)可以找出这些有进位的位，因为进位作用于更高的1位，所以还需要左移1位
所以该加法操作可以分成两部分，XOR的结果（不考虑进位部分）加上(&)<<1（只考虑进位部分）
*/
func getSum(a int, b int) int {
	if b == 0 {
		return a
	} 

	return getSum(a^b, (a&b)<<1)
}

func main() {
	a, b := 1, 2
	fmt.Println(getSum(a,b))

	a, b = 2, 3
	fmt.Println(getSum(a,b))

	a, b = 11, 7
	fmt.Println(getSum(a,b))
}
