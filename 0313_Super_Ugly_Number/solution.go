package main

import (
	"fmt"
)

func min(list, primes, index []int) int {
	min := 2147483647
	for i := range primes {
		if temp := primes[i]*list[index[i]]; temp < min {
			min = temp
		}
	}

	return min
}

func nthSuperUglyNumber(n int, primes []int) int {
	list := make([]int, n)
	list[0] = 1

	index := make([]int, len(primes))
	for i:=1; i<n; i++ {
		next := min(list, primes, index)
		list[i] = next

		for i := range primes {
			if next == primes[i]*list[index[i]] {
				index[i]++
			}
		}
	}

	return list[n-1]
}

func main() {
	nums := []int{12, 1}
	primes := [][]int{[]int{2,7,13,19}, []int{2,3,5}}
	for i, v := range nums {
		fmt.Println(nthSuperUglyNumber(v, primes[i]))
	}
}
