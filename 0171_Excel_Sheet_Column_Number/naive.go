package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {
	rslt, m := 0, 1

	for i := len(columnTitle) - 1; i >= 0 ; i-- {
		rslt += int(columnTitle[i] - 64) * m
		m *= 26
	}	

	return rslt
}

func main() {
	str := []string{"A", "AB", "ZY", "FXSHRXW"}
	for i := range str {
		fmt.Println(titleToNumber(str[i]))
	}
}
