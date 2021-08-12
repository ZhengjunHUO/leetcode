package main

import (
	"fmt"
	"github.com/ZhengjunHUO/godtype"
)

func isValid(curr, min, max *godtype.TreeNode) bool {
	if curr == nil {
		return true
	}

	if min != nil && curr.Val.(int) <= min.Val.(int) {
		return false
	} 
	if max != nil && curr.Val.(int) >= max.Val.(int) {
		return false
	}

	return isValid(curr.Left, min, curr) && isValid(curr.Right, curr, max)
}

func isValidBST(root *godtype.TreeNode) bool {
	return isValid(root, nil, nil)
}

func main() {
	lists := [][]interface{}{[]interface{}{2,1,3}, []interface{}{5,1,4,nil,nil,3,6}, []interface{}{10,5,15,nil,nil,6,20}, []interface{}{5,3,6,2,4,nil,nil,1}}
	for i := range lists {
		fmt.Println(isValidBST(godtype.NewBTree(lists[i])))
	}
}
