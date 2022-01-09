package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func main() {
	a := []int{1, 1, 1, 2, 2, 3, 4, 8, 10}

	// left bound
	x := 2
	i := gtoolkit.FindLeftBound(a, x, func(i, target int) bool{
		return a[i] < target
	})
	fmt.Printf("found %d at index %d in %v\n", x, i, a)

	x = 5
	i = gtoolkit.FindLeftBound(a, x, func (i, target int) bool {
		return a[i] < target
	})
	fmt.Printf("%d not found, can be inserted at index %d in %v\n", x, i, a)

	fmt.Println("========")

	// right bound 
	x = 2
	i = gtoolkit.FindRightBound(a, x, func(i, target int) bool {
		return a[i] > target
	})
	fmt.Printf("found %d at index %d in %v\n", x, i, a)

	x = 5
	i = gtoolkit.FindRightBound(a, x, func(i, target int) bool {
		return a[i] > target
	})
	fmt.Printf("%d not found, right bound at index %d in %v\n", x, i, a)
}
