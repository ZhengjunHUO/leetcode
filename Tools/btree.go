package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func main() {
	godtype.PrintBtree(godtype.NewBTree([]interface{}{3,9,20,nil,nil,15,7}, 0))
}
