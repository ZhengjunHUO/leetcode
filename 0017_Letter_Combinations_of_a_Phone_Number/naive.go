package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func backtrack(digits string, dict map[byte][]byte, rslt *[]string, curr string, currIdx int) {
	if len(curr) == len(digits) {
		*rslt = append(*rslt, curr)
		return
	}

	for _, v := range dict[digits[currIdx]] {
		backtrack(digits, dict, rslt, curr+string(v), currIdx+1)
	} 
}

func letterCombinations(digits string) []string {
	dict := make(map[byte][]byte)
	var base byte = 'a'
	var i byte = '2'
	for ; i<58; i++ {
		if i == 55 || i == 57 {
			dict[i] = []byte{base, base+1, base+2, base+3}
			base += 4
		}else{
			dict[i] = []byte{base, base+1, base+2}
			base += 3
		}
	}

	rslt := make([]string, 0, int(gtoolkit.Factorial(uint64(len(digits)))))
	backtrack(digits, dict, &rslt, "", 0)
	return rslt
}

func main() {
	str := []string{"23", "", "2"}
	for i := range str {
		fmt.Println(letterCombinations(str[i]))
	}
}
