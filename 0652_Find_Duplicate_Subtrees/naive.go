package main

import (
	"github.com/ZhengjunHUO/godtype"
)

func find(curr *godtype.TreeNode, dict map[string]int, rslt *[]*godtype.TreeNode) string {
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

func findDuplicateSubtrees(root *godtype.TreeNode) []*godtype.TreeNode {
	rslt := []*godtype.TreeNode{}
	dict := make(map[string]int)
	find(root, dict, &rslt)
	return rslt
}

func main() {
	list := findDuplicateSubtrees(godtype.NewBTree([]interface{}{1,2,3,4,nil,2,4,nil,nil,nil,nil,4}, 0))
	for _, v := range list {
		godtype.PrintBtree(v)
	}


	list = findDuplicateSubtrees(godtype.NewBTree([]interface{}{2,1,1}, 0))
	for _, v := range list {
		godtype.PrintBtree(v)
	}

	list = findDuplicateSubtrees(godtype.NewBTree([]interface{}{2,2,2,3,nil,3,nil}, 0))
	for _, v := range list {
		godtype.PrintBtree(v)
	}
}
