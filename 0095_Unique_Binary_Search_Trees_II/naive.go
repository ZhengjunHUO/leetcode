package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func generate(l, r int) []*godtype.TreeNode {
	if l > r {
		return []*godtype.TreeNode{nil}
	}

	rslt := []*godtype.TreeNode{}

	for i:=l; i<=r; i++ {
		// 左右子树会返回一个结点串，如果是空还是会返回长度为1的[nil]
		left := generate(l, i-1)
		right := generate(i+1, r)

		for j := range left {
			for k := range right {
				// 排列组合，构成从本结点出发能生成的所有可能
				curr := &godtype.TreeNode{i, left[j], right[k]}	
				rslt = append(rslt, curr)		
			}
		}
	}

	return rslt
}

func generateTrees(n int) []*godtype.TreeNode {
	return generate(1, n)
}

func main() {
	trees := generateTrees(3)	
	for i := range trees {
		godtype.PrintBTreeBFS(trees[i])
	}
}
