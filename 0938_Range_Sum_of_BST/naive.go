package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func rangeSumBST(root *godtype.TreeNode, low int, high int) int {
	rslt := 0
	var dfs func(current *godtype.TreeNode)

	dfs = func(current *godtype.TreeNode) {
		if current == nil {
			return 
		}
	
		curVal := current.Val.(int) 
		if curVal <= low {
			if curVal == low {
				fmt.Printf("[Debug] add %d\n", curVal)
				rslt += curVal
			}
			dfs(current.Right)
		}else if curVal >= high {
			if curVal == high {
				fmt.Printf("[Debug] add %d\n", curVal)
				rslt += curVal
			}
			dfs(current.Left)
		}else{
			fmt.Printf("[Debug] add %d\n", curVal)
			rslt += curVal
			dfs(current.Left)
			dfs(current.Right)
		}
	}

	dfs(root)

	return rslt
}

func main() {
	trees := [][]interface{}{[]interface{}{10,5,15,3,7,nil,18}, []interface{}{10,5,15,3,7,13,18,1,nil,6}}
	lows, highs := []int{7,6}, []int{15,10}
	for i := range trees {
		fmt.Println(rangeSumBST(godtype.NewBTree(trees[i]), lows[i], highs[i]));
	}
}
