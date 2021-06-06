package main

import (
	"fmt"
)

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}


// 从数列构建树
func NewTree(nums []int, index int) *TreeNode {
	// 有子节点
	if 2*index + 2 < len(nums) {
		return &TreeNode{
			Val: nums[index],
			Left: NewTree(nums, 2*index+1),
			Right: NewTree(nums, 2*index+2),
		}
	}
	
	// 为叶子结点
	return &TreeNode{ nums[index], nil, nil, }
}


func rob(root *TreeNode) (rslt int) {
	rslt, _ = robbis(0, 0, root)
	return
}

func robbis(parentValue int, gparentValue int, current *TreeNode) (int, bool) {
	// 为叶子结点
	if current.Left == nil && current.Right == nil {
//		fmt.Printf("Leaf %d\n", current.Val)

		// 叶子其实为nil
		if current.Val == 0 {
			return 0, false
		}

		if temp := current.Val + gparentValue; temp >= parentValue {
//			fmt.Println("return: ", current.Val, true)
			return current.Val, true
		}else{
//			fmt.Println("return: ", 0, false)
			return 0, false
		}
	}

	// 有至少一个子节点
	leftValue, rightValue, leftIsChosen, rightIsChosen := 0, 0, false, false

	if current.Left == nil {
		leftValue, leftIsChosen = 0, false	
	}else{
		leftValue, leftIsChosen = robbis(current.Val, parentValue, current.Left)
	}

	if current.Right == nil {
		rightValue, rightIsChosen = 0, false
	}else{
		rightValue, rightIsChosen = robbis(current.Val, parentValue, current.Right)
	}
	
//	fmt.Printf("Node %d\n", current.Val)
	if leftIsChosen == false && rightIsChosen == false {
//		fmt.Println("Children not chosen, return: ", (current.Val + leftValue + rightValue), true)
		return (current.Val + leftValue + rightValue), true
	}
	
	if leftValue + rightValue + parentValue > current.Val + gparentValue {
//		fmt.Println("Children win, return: ", leftValue + rightValue, false)
		return leftValue + rightValue, false
	}else{
//		fmt.Println("Children lose, return: ", current.Val, true)
		return current.Val, true
	}
}

func main() {
	fmt.Println(rob(NewTree([]int{3,2,3,0,3,0,1}, 0)))
	fmt.Println(rob(NewTree([]int{3,4,5,1,3,0,1}, 0)))
	fmt.Println(rob(NewTree([]int{5,10,3,1,2,7,3}, 0)))
}
