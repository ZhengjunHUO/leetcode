package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

func main() {
	d := datastruct.NewDeque([]int{})
	d.PopFirst()
	d.PopLast()
	d.PushFirst(3)
	d.PushLast(5)
	d.PushFirst(2)
	d.PushLast(10)
	fmt.Println(d.Size())
	fmt.Println(d.PeekFirst())
	fmt.Println(d.PopFirst())
	fmt.Println(d.PopLast())
}
