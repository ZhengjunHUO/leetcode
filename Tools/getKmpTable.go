package main

import (
	"fmt"
	"github.com/ZhengjunHUO/gtoolkit"
)

func main() {
	fmt.Println("[0 1 2 3]")
	fmt.Println(gtoolkit.GetKmpTable("AAAA"))
	fmt.Println()
	
	fmt.Println("[0 0 0 0 0]")
	fmt.Println(gtoolkit.GetKmpTable("ABCDE"))
	fmt.Println()

	fmt.Println("[0 1 0 1 2 0 1 2 3 4 5]")
	fmt.Println(gtoolkit.GetKmpTable("AABAACAABAA"))
	fmt.Println()

	fmt.Println("[0 1 2 0 1 2 3 3 3 4]")
	fmt.Println(gtoolkit.GetKmpTable("AAACAAAAAC"))
	fmt.Println()

	fmt.Println("[0 1 2 0 1 2 3]")
	fmt.Println(gtoolkit.GetKmpTable("AAABAAA"))
	fmt.Println()

	fmt.Println("[0 0 0 1 0 0] ")
	fmt.Println(gtoolkit.GetKmpTable("为我我为人人"))
}
