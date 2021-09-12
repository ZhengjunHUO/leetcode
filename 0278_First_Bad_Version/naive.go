package main

import (
	"fmt"
)

var BAD int 

func isBadVersion(version int) bool {
	return version >= BAD
}

func firstBadVersion(n int) int {
	l, r := 1, n+1

	for l < r {
		m := l + (r-l)/2
		fmt.Printf("try %d, get %v\n", m, isBadVersion(m))
		if isBadVersion(m) == true {
			r = m
		}else{
			l = m + 1
		}
	}

	return l
}

func main() {
	BAD = 4
	fmt.Println(firstBadVersion(5))
}
