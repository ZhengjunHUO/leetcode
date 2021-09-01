package main

import (
	"fmt"
)

const module = 1337

// (a * b) % x = ((a % x)(b % x)) % x
func pow(a, b int) int {
	if b == 0 {
		return 1
	}

	a %= module

	if b%2 == 0 {
		temp := pow(a,b/2)
		return (temp*temp)%module
	}

	return (a * pow(a, b-1))%module
} 

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}

	x := pow(a, b[len(b)-1])
	b = b[:len(b)-1]
	return (x*pow(superPow(a, b), 10))%module
}

func main() {
	a := []int{2, 2, 1, 2147483647}
	b := [][]int{[]int{3}, []int{1,0}, []int{4,3,3,8,5,2}, []int{2,0,0}}

	for i := range a {
		fmt.Println(superPow(a[i], b[i]))
	}
}
