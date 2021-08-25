package main

import (
	"fmt"
	"time"
)

type Node struct {
	Value int
	Children []*Node
	Right, Left, Next *Node
}

func (n *Node) DepthFirstSearch(arr []int) []int {
	arr = append(arr, n.Value)
	for _, child := range n.Children {
		arr = child.DepthFirstSearch(arr)
	}
	return arr
}

// SumPath returns whether there is a path from the root to a leaf node that sums to a value.
func (n *Node) SumPath(sum int) bool {
	if n == nil {
		return false
	}
	if n.Value == sum && n.Left == nil && n.Right == nil {
		return true
	}
	return n.Right.SumPath(sum - n.Value) || n.Left.SumPath(sum - n.Value)
}

// AllSumPath returns all paths in the tree from the node to a leaf that equals a value.
func (n *Node) AllSumPath(sum int) [][]int {
	s := time.Now()
	var p [][]int
	var a []int
	n.recursivePaths(sum, a, &p)
	e := time.Since(s)
	fmt.Printf("recursion1 : %s", e)
	return p
}

// AllSumPath2 returns all paths in the tree from the node to a leaf that equals a value.
func (n *Node) AllSumPath2(sum int) [][]int {
	s := time.Now()
	var p [][]int
	var a []int
	n.recursivePaths2(sum, a, &p)
	e := time.Since(s)
	fmt.Printf("recursion2 : %s", e)
	return p
}

func (n *Node) recursivePaths(sum int, path []int, allPaths *[][]int) {
	if n == nil {
		return
	}

	path = append(path, n.Value)
	if n.Value == sum && n.Left == nil && n.Right == nil {
		*allPaths = append(*allPaths, path)
		path = path[:len(path)-1]
	} else {
		n.Right.recursivePaths(sum - n.Value, path, allPaths)
		n.Left.recursivePaths(sum - n.Value, path, allPaths)
	}
}

func (n *Node) recursivePaths2(sum int, path []int, allPaths *[][]int) {
	if n == nil {
		return
	}

	path = append(path, n.Value)
	if n.Value == sum && n.Left == nil && n.Right == nil {
		*allPaths = append(*allPaths, append([]int{}, path...))
		path = path[:len(path)-1]
	} else {
		n.Right.recursivePaths(sum - n.Value, path, allPaths)
		n.Left.recursivePaths(sum - n.Value, path, allPaths)
	}
}

// SumNumberPath gets the sum of each DFS path.
func (n *Node) SumNumberPath() int {
	s := time.Now()
	var p [][]int
	var a []int
	n.recursivePath(a, &p)
	var sum int
	for _, x := range p {
		sum += sliceToInt(x)
	}
	e := time.Since(s)
	fmt.Printf("SumNumberPath time: %s", e)
	return sum
}

// SumNumberPath2 gets the sum of each DFS path.
func (n *Node) SumNumberPath2() int {
	s := time.Now()
	e := time.Since(s)
	fmt.Printf("SumNumberPath2 time: %s", e)
	return n.rootToLeafPath(0)
}

func (n *Node) rootToLeafPath(sum int) int {
	if n == nil {
		return 0
	}

	sum = 10 * sum + n.Value
	if n.Left == nil && n.Right == nil {
		return sum
	}
	return n.Left.rootToLeafPath(sum) + n.Right.rootToLeafPath(sum)
}

func (n *Node) recursivePath(path []int, allPaths *[][]int) {
	if n == nil {
		return
	}
	path = append(path, n.Value)
	if n.Left == nil && n.Right == nil {
		*allPaths = append(*allPaths, append([]int{}, path...))
		path = path[:len(path)-1]
	} else {
		n.Right.recursivePath(path, allPaths)
		n.Left.recursivePath(path, allPaths)
	}
}

// PathSequence will determine if a given sequence exists from root to leaf.
func (n *Node) PathSequence(seq []int) bool {
	if n == nil {
		return len(seq) == 0
	}
	return n.checkPath(seq, 0)
}

func (n *Node) checkPath(seq []int, idx int) bool {
	if n == nil || n.Value != seq[idx] {
		return false
	}
	if n.Right == nil && n.Left == nil {
		return true
	}
	idx++
	return n.Right.checkPath(seq, idx) || n.Left.checkPath(seq, idx)
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i:=len(s)-1; i>=0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func main() {
	// node := &Node{Value: 3}
	root := &Node{Value: 1}
	root.Left = &Node{Value: 9}
	root.Right = &Node{Value: 7}
	root.Left.Left = &Node{Value: 9}
	root.Right.Left = &Node{Value: 5}
	root.Right.Right = &Node{Value: 4}
	root.Left.Right = &Node{Value: 2}
	// root.Right.Left.Right = &Node{Value: 17}
	fmt.Println(root.SumPath(11))
	fmt.Println(root.AllSumPath(12))
	fmt.Println(root.AllSumPath2(12))
	fmt.Println(root.SumNumberPath())
	fmt.Println(root.SumNumberPath2())
	fmt.Println(root.PathSequence([]int{1, 9, 9}))
}
