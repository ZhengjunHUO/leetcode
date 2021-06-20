package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func main() {
	d := godtype.NewDeque()
	d.PopFirst()
	d.PopLast()
	d.PushFirst(3)
	d.PushLast(5)
	d.PushFirst(2)
	d.PushLast(10)
	fmt.Println(d.Elems)
	fmt.Println(d.Size())
	fmt.Println(d.PeekFirst())
	fmt.Println(d.PopFirst())
	fmt.Println(d.PopLast())
	fmt.Println(d.Elems)
}
