package main

import (
	"fmt"
)

func numTrees(n int) int {
	rslt := 1
	// 按照公式计算C(2n,n)/(n+1)
	// 即(2n)! / (n)!(2n-n)! / (n+1)
	for i:=1; i<=n; i++ {
		rslt = rslt*(i+n)/i
	}
	
	return rslt/(n+1)
}

func main() {
	ns := []int{3,1,6,7}
	for _, v := range ns {
		fmt.Println(numTrees(v))
	}
}
