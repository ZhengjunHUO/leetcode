package main

import (
	"fmt"	
)

func countPairs(deliciousness []int) int {
	mod := 1000000007
	dict := make(map[int]int)
	rslt := 0
	for _,d := range deliciousness {
		target := 1
		for j:=0;j<22;j++ {
			if v, ok := dict[target-d]; ok {
				rslt += v		
			}
			target *= 2
		} 	
		dict[d]+=1
	} 
	return rslt%mod
}

func main() {
	d1 := []int{1,3,5,7,9}
	d2 := []int{1,1,1,3,3,3,7}
	fmt.Println(countPairs(d1))
	fmt.Println(countPairs(d2))
}
