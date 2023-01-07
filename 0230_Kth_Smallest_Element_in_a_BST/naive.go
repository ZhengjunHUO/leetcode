package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

func kthSmallest(root *graph.TreeNode, k int) int {
	return find(root, &k)
}

func find(root *graph.TreeNode, k *int) int {
	if root == nil {
		// 返回的值无所谓
		return 0
	}

	rslt := find(root.Left, k)
	if *k == 0 {
		// 如果左子树找到了答案会沿着左侧直接返回到最上层
		return rslt
	}

	*k -= 1
	if *k == 0 {
		return root.Val.(int)
	}

	return find(root.Right, k)
}

func main() {
	trees := [][]interface{}{[]interface{}{3,1,4,nil,2}, []interface{}{5,3,6,2,4,nil,nil,1}}
	ks := []int{1, 3}
	for i := range trees {
		fmt.Println(kthSmallest(graph.NewBTree(trees[i]), ks[i]))
	}
}
