package main

import (
	"fmt"
)

// 参考@san89kalp的方法
func hasAlternatingBits(n int) bool {
	// 二进制表示如果是0和1相间的话，位移一位后XOR可以得到全1
	// 如101 ^ 010 => 111 
	n = n ^ (n >> 1)
	// 二进制为全1的话再加1会进位一个1后面变为全0，与原数相与可得到0
	// 如111 & 1000 => 0
	return (n & (n+1)) == 0
}

func main() {
	ns := []int{5,7,11,10,3}
	for i := range ns {
		fmt.Println(hasAlternatingBits(ns[i]))
	}
}
