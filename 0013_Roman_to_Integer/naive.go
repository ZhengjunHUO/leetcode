package main

import (
	"fmt"
)

func romanToInt(s string) int {
	dict := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	rslt := 0

	for i := 0; i < len(s) - 1; i++ {
		if dict[s[i]] < dict[s[i+1]] {
			rslt -= dict[s[i]]
		}else{
			rslt += dict[s[i]]
		}
	}

	return rslt + dict[s[len(s)-1]]
}

func main() {
	test := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for _, v := range test {
		fmt.Println(romanToInt(v))
	}
}
