package main

import (
	"fmt"
)

func convertToTitle(columnNumber int) string {
	rslt := ""
	
	for columnNumber > 0 {
		temp, i := columnNumber, 0 
		for temp/27 > 0 {
			temp /= 26
			i++
		}

		rslt += string(byte(64+temp))

		for i > 0 {
			temp *= 26
			i--
		}

		columnNumber -= temp
	}

	return rslt
}

func main() {
	cn := []int{1, 28, 701, 2147483647}
	for i := range cn {
		fmt.Println(convertToTitle(cn[i]))
	}
}
