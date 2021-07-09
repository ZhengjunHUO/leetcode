package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func main() {
	s := []int{1,2,3,4,5}
	fmt.Println(gtoolkit.IntSliceToString(s))
	s1 := gtoolkit.IntSliceRemove(s, 2)
	fmt.Println(gtoolkit.IntSliceToString(s))
	fmt.Println(gtoolkit.IntSliceToString(s1))

	sl := []int{10,20,30,40,50}
	fmt.Println(gtoolkit.IntSliceToString(sl))
	gtoolkit.IntSliceRemoveInplace(&sl, 3)
	fmt.Println(gtoolkit.IntSliceToString(sl))
}
