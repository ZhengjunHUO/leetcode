package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

// 有一些recipes在supplies的基础上还基于其他recipes，所以需要知道其依赖关系先处理没有依赖的receipes
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// 针对recipes构建dag
	numV := len(recipes)
	dag := graph.NewDag(numV)

	// 将recipes做成字典便于根据名字查询index
	dictRecp := make(map[string]int)
	for i := range recipes {
		dictRecp[recipes[i]] = i
	}
	//fmt.Println(dictRecp)

	// 将原料做成字典，后续也会将完成的recipes加入
	dictSupp := make(map[string]struct{})
	for i := range supplies {
		dictSupp[supplies[i]] = struct{}{}
	}
	//fmt.Println(dictSupp)

	// 分析ingredients，探明不同recipes之间的依赖关系，作为一条边加入dag
	for i:=0; i<numV; i++ {
		for j:=0; j<len(ingredients[i]); j++ {
			if v, ok := dictRecp[ingredients[i][j]]; ok {
				dag.AddEdge(v, i)
			}
		}
	}
	// 获得所有边的信息后拓扑排序
	dag.TopologicalSort()
	//fmt.Println(dag.Sorted.Elems)

	ret := []string{}
	// 依次pop出无依赖性的recipes到强依赖性的recipes
	for dag.Sorted.Size() > 0 {
		idx := dag.Sorted.Pop()
		isPossible := true
		for i := range ingredients[idx] {
			// 缺少原料，不能做出
			if _, ok := dictSupp[ingredients[idx][i]]; !ok {
				isPossible = false
				break
			}
		}

		if isPossible {
			// 将完成的recipes加入原料字典
			dictSupp[recipes[idx]] = struct{}{}
			// 加入返回值中
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
