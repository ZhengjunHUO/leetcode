package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func main() {
	s := godtype.NewQueue()
	fmt.Println(s.Pop())

	s.Push(3)
	s.Push(6)
	s.Push(9)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
}
