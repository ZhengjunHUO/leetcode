package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func main() {
	g := godtype.NewDag(6)
	g.AddEdge(5, 2)
        g.AddEdge(5, 0)
        g.AddEdge(4, 0)
        g.AddEdge(4, 1)
        g.AddEdge(2, 3)
        g.AddEdge(3, 1)
	g.TopologicalSort()

	fmt.Println(g.Sorted.Elems)
	for g.Sorted.Size() > 0 {
		fmt.Println(g.Sorted.Pop().(int))
	}
}
