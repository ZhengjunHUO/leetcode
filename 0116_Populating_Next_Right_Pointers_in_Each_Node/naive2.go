package main

import (
	"fmt"
)

type Node struct {
	Val  interface{}
	Left *Node
	Right *Node
	Next *Node
}

func buildTree(elems []interface{}, index int) *Node {
	if len(elems) == 0 {
		return nil
	}

        // 为叶子结点
        if 2*index + 2 > len(elems) {
		return &Node{ elems[index], nil, nil, nil}
	}

        // 有子节点
	var l, r *Node
        if 2*index + 1 < len(elems) && elems[2*index+1] != nil {
		l = buildTree(elems, 2*index+1)
	}
	
        if 2*index + 2 < len(elems) && elems[2*index+2] != nil {
		r = buildTree(elems, 2*index+2)
	}
	
	return &Node{
		Val: elems[index],
		Left: l,
		Right: r,
		Next: nil,
        }
}

var count int
func printIndent(n int) {
    for i := 0; i < n; i++ {
        fmt.Printf("  ")
    }
}

func printTree(root *Node) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	printIndent(count)
	if root.Next != nil {
		fmt.Printf("Current node's value: %v; next: %v\n", root.Val, root.Next.Val)
	}else{
		fmt.Printf("Current node's value: %v; next: %v\n", root.Val, root.Next)
	}
	if root.Left !=	nil {
		printIndent(count)
		fmt.Printf("%v have a left child\n", root.Val)
		count++
		printTree(root.Left)
		count--
	}
	if root.Right != nil {
		printIndent(count)
		fmt.Printf("%v have a right child\n", root.Val)
		count++
		printTree(root.Right)
		count--
	}
}

func connectbis(left *Node, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right

	connectbis(left.Left, left.Right)
	connectbis(right.Left, right.Right)
	connectbis(left.Right, right.Left)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectbis(root.Left, root.Right)
	return root
}

func main() {
	tree := []interface{}{1,2,3,4,5,6,7}
	printTree(connect(buildTree(tree, 0)))
}
