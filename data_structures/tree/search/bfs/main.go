package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Children []*Node
	Right, Left *Node
}

func (n *Node) BreadthFirstSearch(arr []int) {
	queue := []*Node{n}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		arr = append(arr, cur.Value)
		for _, child := range n.Children {
			queue = append(queue, child)
		}
	}
}

// TraversePrint will print each level of the tree as a slice.
func (n *Node) TraversePrint() [][]int {
	var res [][]int
	if n == nil {
		return res
	}
	s := []*Node{n}
	for len(s) > 0 {
		level := len(s)
		var arr []int
		for i:=0; i<level; i++ {
			node := s[0]
			arr = append(arr, node.Value)
			s = s[1:]
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
		res = append(res, arr)
	}
	return res
}

// TraversePrintReverse will print each level of the tree as a slice.
func (n *Node) TraversePrintReverse() [][]int {
	var res [][]int
	if n == nil {
		return res
	}
	s := []*Node{n}
	for len(s) > 0 {
		level := len(s)
		var arr []int
		for i:=0; i<level; i++ {
			node := s[0]
			arr = append(arr, node.Value)
			s = s[1:]
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
		res = append([][]int{arr}, res...)
	}
	return res
}

func (n *Node) traverse() {
	if n != nil {
		s := []*Node{n}
		for len(s) > 0 {
			node := s[0]
			fmt.Printf(strconv.Itoa(node.Value) + " ")
			s = s[1:]
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
	}
}

func (n *Node) postorder() {
	if n != nil {
		n.Left.postorder()
		n.Right.postorder()
		fmt.Println(strconv.Itoa(n.Value) + " ")
	}
}

func main() {
		root := &Node{Value: 12}
		root.Left = &Node{Value: 7}
		root.Right = &Node{Value: 1}
		root.Left.Left = &Node{Value: 9}
		root.Right.Left = &Node{Value: 10}
		root.Right.Right = &Node{Value: 5}
		fmt.Println(root.TraversePrint())
		fmt.Println(root.TraversePrintReverse())
		root.traverse()
		root.postorder()
}
