package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func main() {
	uf := graph.NewUF(10)
	fmt.Println(uf.Count(), uf.Parent(), uf.Size())

	uf.Union(0,1)
	uf.Union(0,6)
	uf.Union(2,3)
	uf.Union(2,5)
	uf.Union(1,3)
	uf.FindRoot(3)
	c, p, s := uf.Count(), uf.Parent(), uf.Size()
	fmt.Println(uf.Count(), uf.Parent(), uf.Size())
	fmt.Println("Saved")

	uf.Union(6,5)
	fmt.Println(uf.Count(), uf.Parent(), uf.Size())

	fmt.Println(uf.IsLinked(3,5))
	fmt.Println(uf.Count(), uf.Parent(), uf.Size())

	fmt.Println(uf.FindRoot(7))

	uf.SetParent(p)
	uf.SetSize(s)
	uf.SetCount(c)
	fmt.Println("Loading ... ")
	fmt.Println(uf.Count(), uf.Parent(), uf.Size())

	uf.Union(7,8)
	fmt.Println(uf.Count(), uf.Parent(), uf.Size())
	fmt.Println(c, p, s)
}
