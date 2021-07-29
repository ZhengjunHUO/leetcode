package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n < 0 {
		return float64(1.0) / myPow(x, -n)
	}

	if n == 0 {
		return 1
	}
	
	temp := myPow(x, n/2)
	if n%2 == 0 {
		return temp*temp
	}
	return x*temp*temp
}

func main() {
	xs := []float64{2.00000, 2.10000, 2.00000} 
	ns := []int{10, 3, -2}

	for i := range xs {
		fmt.Println(myPow(xs[i], ns[i]))
	}
}
