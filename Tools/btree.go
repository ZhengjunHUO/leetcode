package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
)

func main() {
	tree := graph.NewBTree([]interface{}{3,9,20,nil,nil,15,7})
	graph.PrintBTreeDFS(os.Stdout, tree)
	graph.PrintBTreeBFS(os.Stdout, tree)
}
