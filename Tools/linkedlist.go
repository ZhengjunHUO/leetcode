package main

import (
	"github.com/ZhengjunHUO/godtype"
	"fmt"
)

func main() {
	godtype.PrintList(godtype.NewList([]string{"foo", "bar", "huo"}))

	l := godtype.NewList([]int{7,2,4,3})
	godtype.PrintList(l)
	if l.Val.(int) < 10 {
		fmt.Println(l.Val)
	}

	godtype.NewList(10)
	godtype.NewList("hello")
}
