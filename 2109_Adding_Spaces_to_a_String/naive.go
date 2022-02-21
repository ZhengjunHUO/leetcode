package main

import (
	"fmt"
)

func addSpaces(s string, spaces []int) string {
	var ret string
	var prev int
	for _,v := range spaces {
		ret += s[prev:v]+" "
		prev = v
	}

	ret += s[prev:]

	return ret
}

func main() {
	s := []string{"LeetcodeHelpsMeLearn", "icodeinpython", "spacing"}
	spaces := [][]int{[]int{8,13,15}, []int{1,5,7,9}, []int{0,1,2,3,4,5,6}}

	for i := range s {
		fmt.Println(addSpaces(s[i], spaces[i]))
	}
}
