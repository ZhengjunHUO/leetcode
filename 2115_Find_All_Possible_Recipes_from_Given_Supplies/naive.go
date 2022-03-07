package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	numV := len(recipes)
	dag := godtype.NewDag(numV)

	dictRecp := make(map[string]int)
	for i := range recipes {
		dictRecp[recipes[i]] = i
	}
	//fmt.Println(dictRecp)

	dictSupp := make(map[string]struct{})
	for i := range supplies {
		dictSupp[supplies[i]] = struct{}{}
	}
	//fmt.Println(dictSupp)

	for i:=0; i<numV; i++ {
		for j:=0; j<len(ingredients[i]); j++ {
			if v, ok := dictRecp[ingredients[i][j]]; ok {
				dag.AddEdge(v, i)
			}
		}
	}
	dag.TopologicalSort()
	//fmt.Println(dag.Sorted.Elems)

	ret := []string{}
	for dag.Sorted.Size() > 0 {
		idx := dag.Sorted.Pop().(int)
		isPossible := true
		for i := range ingredients[idx] {
			if _, ok := dictSupp[ingredients[idx][i]]; !ok {
				isPossible = false
				break
			}
		}

		if isPossible {
			dictSupp[recipes[idx]] = struct{}{}
			ret = append(ret, recipes[idx])
		}
	}

	return ret
}

func main() {
	recipes := [][]string{[]string{"bread"}, []string{"bread","sandwich"}, []string{"bread","sandwich","burger"}}
	ingredients := [][][]string{[][]string{[]string{"yeast","flour"}}, [][]string{[]string{"yeast","flour"},[]string{"bread","meat"}}, [][]string{[]string{"yeast","flour"},[]string{"bread","meat"},[]string{"sandwich","meat","bread"}}}
	supplies := [][]string{[]string{"yeast","flour","corn"}, []string{"yeast","flour","meat"}, []string{"yeast","flour","meat"}}

	for i := range recipes {
		fmt.Println(findAllRecipes(recipes[i], ingredients[i], supplies[i]))
	}
}
