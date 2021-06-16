package main

import (
	"fmt"
)

func isNumber(s string) bool {
	valid := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'} 	
	dict := make(map[byte]int)

	for _, b := range valid {
		dict[b] = 1
	}

	i := 0

	if s[i] == '-' || s[i] == '+' {
		i += 1
	}

	if _, ok := dict[s[i]]; ok {
		i += 1
		dict['e'], dict['E'] = 1, 1
	}else{
		return false
	}

	for i < len(s) {
		if _, ok := dict[s[i]]; ok {
			if s[i] == '.' {
				delete(dict, '.')
			}

			if s[i] == 'e' || s[i] == 'E' {
				if i + 1 == len(s) {
					return false
				}

				if s[i+1] == '+' || s[i+1] == '-' {
					i += 1
				}
				delete(dict, 'e')
				delete(dict, 'E')
				delete(dict, '.')
			}
		}else{
			return false
		}

		i += 1
	}	

	return true    
}

func main() {
	str := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789", "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"}
	for _, s := range str {
		fmt.Printf("%s \tis string ? %v\n", s, isNumber(s))
	}
}
