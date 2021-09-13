package main

import (
	"fmt"
)

func isPowerOfFour(n int) bool {
	return n > 0 && ( n&(n-1) == 0 ) && ( n&0x55555555 != 0 )
}

func main() {
	ns := []int{16,5,1}
	for i := range ns {
		fmt.Println(isPowerOfFour(ns[i]))
	}
}
