package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func main() {
	s := []int{34,-10,23,1,24,-9,75,33,54,8}
	fmt.Println(s)
	gtoolkit.SortInts(s)	
	fmt.Println(s)
}
