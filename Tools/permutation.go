package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func main() {
	s := [][]int{[]int{1,2,3}, []int{1,2,3,4}}
	for i := range s {
		temp := gtoolkit.Permutation(s[i])
		fmt.Println(len(temp), temp)
	}
}
