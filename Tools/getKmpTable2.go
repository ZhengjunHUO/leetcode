package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func main() {
	pats := []string{"AABA", "ABABCABAB", "AAAAB", "ABABAC", "AAAA", "小猫咪"}
	txts := []string{"AABAACAADAABAABA", "ABABDABACDABABCABAB", "AAAAAAAAAAAAAAAAAB", "ABABABCABABABCABABABC", "AAAAABAAABA", "福福是一只小猫咪"}

	/*
		[0 9 12]
		[10]
		[13]
		[]
		[0 1]
	*/
	for i := range pats {
		fmt.Println(godtype.NewPatternFinder(pats[i]).FindIn(txts[i]))
	}
}
