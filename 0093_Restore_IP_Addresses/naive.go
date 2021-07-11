package main

import (
	"fmt"
)

func backtrack(s string, rslt *[]string, curr string, count int, currIdx int) {
	if currIdx == len(s) {
		if count == 4 {
			*rslt = append(*rslt, curr[:len(curr)-1])
		}
		return
	}

	// 1 digit [0-9]
	backtrack(s, rslt, curr+string(s[currIdx])+".", count+1, currIdx+1)

	if s[currIdx] != byte('0') && currIdx+1 < len(s) {
		// 2 digits [10-99]
		backtrack(s, rslt, curr+s[currIdx:currIdx+2]+".", count+1, currIdx+2)
		if currIdx+2 < len(s) && int(s[currIdx]-48)*100+int(s[currIdx+1]-48)*10+int(s[currIdx+2]-48) <= 255 {
			// 3 digits [100-255]
			backtrack(s, rslt, curr+s[currIdx:currIdx+3]+".", count+1, currIdx+3)
		}
	} 
}

func restoreIpAddresses(s string) []string {
	rslt := []string{}    
	backtrack(s, &rslt, "", 0, 0)
	return rslt
}

func main() {
	ss := []string{"25525511135", "0000", "1111", "010010", "101023"}
	for _, s := range ss {
		fmt.Println(restoreIpAddresses(s))
	}
}
