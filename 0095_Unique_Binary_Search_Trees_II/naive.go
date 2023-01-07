package main

import (
	"os"
	"github.com/ZhengjunHUO/goutil/graph"
)

func generate(l, r int) []*graph.TreeNode {
	if l > r {
		return []*graph.TreeNode{nil}
	}

	rslt := []*graph.TreeNode{}

	for i:=l; i<=r; i++ {
		// 左右子树会返回一个结点串，如果是空还是会返回长度为1的[nil]
		left := generate(l, i-1)
		right := generate(i+1, r)

		for j := range left {
			for k := range right {
				// 排列组合，构成从本结点出发能生成的所有可能
				curr := &graph.TreeNode{i, left[j], right[k]}
				rslt = append(rslt, curr)
			}
		}
	}

	return rslt
}

func generateTrees(n int) []*graph.TreeNode {
	return generate(1, n)
}

func main() {
	trees := generateTrees(3)
	for i := range trees {
		graph.PrintBTreeBFS(os.Stdout, trees[i])
	}
}
