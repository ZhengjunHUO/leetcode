package main

import (
	"fmt"
)

/*
from "zestypanda":

The solution is based on 2 facts:
1) the number of length k string without consecutive 1 is Fibonacci sequence f(k);
For example, if k = 5, the range is 00000-11111. We can consider it as two ranges, which are 00000-01111 and 10000-10111. Any number >= 11000 is not allowed due to consecutive 1. The first case is actually f(4), and the second case is f(3), so f(5)= f(4)+f(3).
2) Scan the number from most significant digit, i.e. left to right, in binary format. If we find a '1' with k digits to the right, count increases by f(k) because we can put a '0' at this digit and any valid length k string behind; After that, we continue the loop to consider the remaining cases, i.e., we put a '1' at this digit. If consecutive 1s are found, we exit the loop and return the answer. By the end of the loop, we return count+1 to include the number n itself.
For example, if n is 10010110,
we find first '1' at 7 digits to the right, we add range 00000000-01111111, which is f(7);
second '1' at 4 digits to the right, add range 10000000-10001111, f(4);
third '1' at 2 digits to the right, add range 10010000-10010011, f(2);
fourth '1' at 1 digits to the right, add range 10010100-10010101, f(1);
Those ranges are continuous from 00000000 to 10010101. And any greater number <= n will have consecutive 1.
*/

func findIntegers(n int) int {
	// n <= 10^9, 十六进制下为0x3B9ACA00, 最大需要30位表示
	dict := make([]int, 30)
	dict[0], dict[1] = 1, 2
	for i:=2; i<30; i++ {
		dict[i] = dict[i-2] + dict[i-1]
	} 

	rslt, prev, p := 0, 0, 29
	for p >= 0{
		if (n & (1<<p)) > 0 {
			rslt += dict[p]
			// 已经有两个连续的1，没必要继续了
			if prev == 1 {
				return rslt
			}
			prev = 1
		}else{
			prev = 0
		}
		p -= 1
	}
	
	// 如果经历完一遍for没有中途return的话，说明self也是符合条件的，应该加入rslt
	return rslt+1
}

func main() {
//	fmt.Println(findIntegers(5))
//	fmt.Println(findIntegers(1))
//	fmt.Println(findIntegers(2))
	fmt.Println(findIntegers(16))
}
