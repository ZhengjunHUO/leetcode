package main

import (
	"fmt"
	"math"
)

// char in t is unique 
func minWindow(s string, t string) string {
	primes := sieveOfEratosthenes(len(t)*len(t))
	
	// [dict] key: ASCII ; value: prime number
	dict := make(map[int]int)
	tprod := 1
	for i := range t {
		dict[int(t[i])] = primes[i]
		tprod *= primes[i]
	} 

	// [cumprod] key: production cumulative; value: index 
	cumprod := make(map[int]int)
	cumprod[1] = -1
	prevKey := 1

	prod := 1
	size := 100000
	startPoint := -1

	for i := range s {
		if v, ok := dict[int(s[i])]; ok {
			prod *= v
			if prod >= tprod {
				prod /= tprod 
			} 

			cumprod[prevKey] = i
			prevKey = prod

			if val, exist := cumprod[prod]; exist {
				temp := i - val + 1	
				if temp < size {
					size = temp
					startPoint = val
				}
			}
		}
	}

	if size < 100000 {
		return s[startPoint:startPoint+size]
	}   

	return ""
}

func sieveOfEratosthenes(n int) []int {
	isNotPrime := make([]bool, n+1)
	prime := []int{}
	pivot := int(math.Ceil(math.Sqrt(float64(n))))

	for i:=2; i<n+1; i++ {
		if isNotPrime[i] == false {
			prime = append(prime, i)
			if i<=pivot {
				for j:=i*i; j<=n; j+=i {
					isNotPrime[j] = true
				}
			}
		}
	}

	return prime
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s,t))
}
