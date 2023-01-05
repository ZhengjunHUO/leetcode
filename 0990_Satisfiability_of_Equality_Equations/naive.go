package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func equationsPossible(equations []string) bool {
	uf := graph.NewUF(26)

	for i := range equations {
		if equations[i][1] == '=' {
			uf.Union(int(equations[i][0] - 'a'), int(equations[i][3] - 'a'))
		}
	}

	for i := range equations {
		if equations[i][1] == '!' {
			if uf.IsLinked(int(equations[i][0] - 'a'), int(equations[i][3] - 'a')) {
				return false
			}
		}
	}

	return true
}

func main() {
	eqs := [][]string{[]string{"a==b","b!=a"}, []string{"b==a","a==b"}, []string{"a==b","b==c","a==c"}, []string{"a==b","b!=c","c==a"}, []string{"c==c","b==d","x!=z"}}
	for i := range eqs {
		fmt.Println(equationsPossible(eqs[i]))
	}
}
