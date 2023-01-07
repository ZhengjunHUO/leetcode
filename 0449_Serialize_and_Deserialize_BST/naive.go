package main

import (
	"fmt"
	"strings"
	"github.com/ZhengjunHUO/goutil/graph"
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *graph.TreeNode) string {
	if root == nil {
		return ""
	}

	rslt := ""
	q := datastruct.NewQueue([]*graph.TreeNode{})
	q.Push(root)

	var emptyNode graph.TreeNode
	size := 0

	loop: for !q.IsEmpty() {
		size = q.Size()
		emptyNum := 0
		for i:=0; i<size; i++ {
			node := q.Pop()
			if *node != emptyNode {
				rslt = rslt + node.Val.(string) + ","
			}else{
				emptyNum++
				rslt = rslt + "nil,"
			}

			if emptyNum == size {
				break loop
			}

			if node.Left != nil {
				q.Push(node.Left)
			}else{
				q.Push(&graph.TreeNode{})
			}

			if node.Right != nil {
				q.Push(node.Right)
			}else{
				q.Push(&graph.TreeNode{})
			}
		}
	}

	return rslt[:len(rslt)-size*4-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *graph.TreeNode {
	temp := strings.Split(data, ",")
	itf := make([]interface{}, len(temp))
	for i := range temp {
		if temp[i] == "nil" {
			itf[i] = nil
		}else{
			itf[i] = temp[i]
		}
	}

	return graph.NewBTree(itf)
}

func main() {
	//s := "2,1,3"
	s := "0,-10,5,nil,-3,nil,9"
	obj := Constructor()

	if s == obj.serialize(obj.deserialize(s)) {
		fmt.Println("ok")
	}else{
		fmt.Println("ko")
	}
}
