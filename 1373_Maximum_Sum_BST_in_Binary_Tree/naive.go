package main

import (
	"fmt"
	"github.com/ZhengjunHUO/goutil/graph"
)

const minint = int(^uint(0) >> 1)
var maxsum int

func recursive(curr, min, max *graph.TreeNode) int {
	if curr == nil {
		return 0
	}

	// 叶子结点，考虑其可能自成一BST树，需要把它的值和maxsum比较
	// 并判断是否包含该叶子的树是否符合BST条件，来决定返回的值
	if curr.Left == nil && curr.Right == nil {
		if curr.Val.(int) > maxsum {
			maxsum = curr.Val.(int)
		}

		if (min != nil && curr.Val.(int) < min.Val.(int)) || ( max != nil && curr.Val.(int) > max.Val.(int)) {
			return minint
		}
		return curr.Val.(int)
	}

	// 如果当前结点违反了BST规则，返回int的最小值
	if (min != nil && curr.Val.(int) < min.Val.(int)) || ( max != nil && curr.Val.(int) > max.Val.(int)) {
		return minint
	}

	left := recursive(curr.Left, min, curr)
	right := recursive(curr.Right, curr, max)

	// 后序
	// 如果子树不符合BST定义，则以当前结点为根的树也不符合，并递归传递上去
	if left == minint || right == minint {
		return minint
	}

	// 如果当前子树符合BST定义，则尝试更新maxsum值，并将和返回上一级
	temp := curr.Val.(int) + left + right
	if temp > maxsum {
		maxsum = temp
	}

	return temp
}

func maxSumBST(root *graph.TreeNode) int {
	maxsum = -minint-1
	recursive(root, nil, nil)
	if maxsum == -minint-1 {
		return 0
	}

	return maxsum
}

func main() {
	trees := [][]interface{}{[]interface{}{1,4,3,2,4,2,5,nil,nil,nil,nil,nil,nil,4,6}, []interface{}{4,3,nil,1,2}, []interface{}{}, []interface{}{2,1,3}, []interface{}{5,4,8,3,nil,6,3}}
	for i := range trees {
		fmt.Println(maxSumBST(graph.NewBTree(trees[i])))
	}
}
