package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	tab := make([][]string, 4)
	tab[0] = []string{"I", "V"}
	tab[1] = []string{"X", "L"}
	tab[2] = []string{"C", "D"}
	tab[3] = []string{"M"}

	rslt := ""
	lvl := 0

	for num > 0 {
		curr := num%10
		num /= 10

		switch {
		case curr < 4:
			rslt = strings.Repeat(tab[lvl][0], curr) + rslt
		case curr == 4:
			rslt = tab[lvl][0] + tab[lvl][1] + rslt
		case curr > 4 && curr < 9:
			rslt = tab[lvl][1] + strings.Repeat(tab[lvl][0], curr - 5) + rslt
		case curr == 9:
			rslt = tab[lvl][0] + tab[lvl+1][0] + rslt
		}

		lvl += 1

	}

	return rslt
}

func main() {
	test := []int{3,4,9,58,1994}
	for _, v := range test {
		fmt.Println(intToRoman(v))
	}
}
