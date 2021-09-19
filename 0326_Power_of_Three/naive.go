package main

import (
	"fmt"
)

// 2^31 - 1内最大的3^x为3^19 = 1162261467
// 对于其他质数也可以使用类似方法
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func main() {
	n := []int{27,0,9,45}
	for i := range n {
		fmt.Println(isPowerOfThree(n[i]))
	}
}
