package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func main() {
	s := "abcde"
	fmt.Printf("%s reversed: %s\n", s, gtoolkit.ReverseStr(s))

	s = "你好,世界!"
	fmt.Printf("%s reversed: %s\n", s, gtoolkit.ReverseStr(s))
}
