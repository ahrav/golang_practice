package main

type Node struct {
	Value int
	Children []*Node
}

func (n *Node) DepthFirstSearch(arr []int) []int {
	arr = append(arr, n.Value)
	for _, child := range n.Children {
		arr = child.DepthFirstSearch(arr)
	}
	return arr
}
