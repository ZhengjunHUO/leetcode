package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func countPrimes(n int) int {
	return len(gtoolkit.SieveOfEratosthenes(n))
}

func main() {
	nums := []int{10, 0, 1}
	for _, v := range nums {
		fmt.Println(countPrimes(v))
	}
}
