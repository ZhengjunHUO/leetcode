package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func main() {
	tree := godtype.NewBTree([]interface{}{3,9,20,nil,nil,15,7})
	godtype.PrintBTreeDFS(tree)
	godtype.PrintBTreeBFS(tree)
}
