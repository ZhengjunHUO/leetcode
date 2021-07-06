package main

import (
	"fmt"
)

func partition(s string) [][]string {
	var rslt [][]string
	backtrack(s, &rslt, []string{}, 0)
    	return rslt
}

func backtrack(s string, rslt *[][]string, curr []string, currIdx int) {
	if currIdx == len(s) {
		valid := make([]string, len(curr), cap(curr))
		copy(valid, curr)
		*rslt = append(*rslt, valid)
	}else{
		for i := currIdx; i < len(s); i++ {
			if isValid(s, currIdx, i) {
				curr = append(curr, s[currIdx:i+1])
				backtrack(s, rslt, curr, i+1)
				curr = curr[0:len(curr)-1]
			}
		}

	}
}

func isValid(pal string, lp int, rp int) bool {
	for lp < rp {
		if pal[lp] == pal[rp] {
			lp += 1
			rp -= 1
		}else{
			return false
		}	
	}

	return true
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
}
