package main

type Node struct {
	Value int
	Children []*Node
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
