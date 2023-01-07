package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
)

func find(curr *graph.TreeNode, dict map[string]int, rslt *[]*graph.TreeNode) string {
	if curr == nil {
		return "nil"
	}

	str := string(curr.Val.(int)) + "," + find(curr.Left, dict, rslt) + "," + find(curr.Right, dict, rslt)
	if v, ok := dict[str]; ok {
		if v == 1 {
			*rslt = append(*rslt, curr)
		}
	}
	dict[str] += 1

	return str
}

func findDuplicateSubtrees(root *graph.TreeNode) []*graph.TreeNode {
	rslt := []*graph.TreeNode{}
	dict := make(map[string]int)
	find(root, dict, &rslt)
	return rslt
}

func main() {
	list := findDuplicateSubtrees(graph.NewBTree([]interface{}{1,2,3,4,nil,2,4,nil,nil,nil,nil,4}))
	for _, v := range list {
		graph.PrintBTreeDFS(os.Stdout, v)
	}

	list = findDuplicateSubtrees(graph.NewBTree([]interface{}{2,1,1}))
	for _, v := range list {
		graph.PrintBTreeDFS(os.Stdout, v)
	}

	list = findDuplicateSubtrees(graph.NewBTree([]interface{}{2,2,2,3,nil,3,nil}))
	for _, v := range list {
		graph.PrintBTreeDFS(os.Stdout, v)
	}
}
