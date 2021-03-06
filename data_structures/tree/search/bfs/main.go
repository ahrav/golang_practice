package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/ahrav/learngo/algos/utils"
)

type Node struct {
	Value             int
	Children          []*Node
	Right, Left, Next *Node
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
		levelSize := len(s)
		var arr []int
		for i := 0; i < levelSize; i++ {
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

// ConnectLevelOrder connects each node to its level order successor.
func (n *Node) ConnectLevelOrder() {
	if n == nil {
		return
	}
	s := []*Node{n}
	for len(s) > 0 {
		levelSize := len(s)
		for i := 0; i < levelSize; i++ {
			node := s[0]
			if i == levelSize-1 {
				node.Next = nil
			} else {
				node.Next = s[1]
			}
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

// ConnectAllLevelOrder connects each node to its level order sibling.
func (n *Node) ConnectAllLevelOrder() {
	if n == nil {
		return
	}
	s := []*Node{n}
	for len(s) > 0 {
		node := s[0]
		s = s[1:]
		if node.Right != nil {
			s = append(s, node.Right)
		}
		if node.Left != nil {
			s = append(s, node.Left)
		}
		if len(s) > 0 {
			node.Next = s[0]
		} else {
			node.Next = nil
		}
	}
}

// RightView returns all nodes that are the last node on each level.
func (n *Node) RightView() []Node {
	if n == nil {
		return nil
	}
	var res []Node
	s := []*Node{n}
	for len(s) > 0 {
		levelSize := len(s)
		for i := 0; i < levelSize; i++ {
			node := s[0]
			if i == levelSize-1 {
				res = append(res, *node)
			}
			s = s[1:]
			if node.Right != nil {
				s = append(s, node.Right)
			}
			if node.Left != nil {
				s = append(s, node.Left)
			}
		}
	}
	return res
}

// TraverseAvgPrint will print each level's average.
func (n *Node) TraverseAvgPrint() []float64 {
	var res []float64
	if n == nil {
		return res
	}
	s := []*Node{n}
	for len(s) > 0 {
		levelSize := len(s)
		var sum int
		for i := 0; i < levelSize; i++ {
			node := s[0]
			sum += node.Value
			s = s[1:]
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(levelSize))
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
		levelSize := len(s)
		var arr []int
		for i := 0; i < levelSize; i++ {
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

// TraverseZigZagPrint will print each level of the tree as a slice.
func (n *Node) TraverseZigZagPrint() [][]int {
	var res [][]int
	if n == nil {
		return res
	}
	s := []*Node{n}
	for len(s) > 0 {
		levelSize := len(s)
		var arr []int
		for i := 0; i < levelSize; i++ {
			node := s[0]
			if levelSize%2 == 0 {
				arr = append([]int{node.Value}, arr...)
			} else {
				arr = append(arr, node.Value)
			}
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

// MinDepth will return shortest depth of the tree.
func (n *Node) MinDepth() int {
	if n == nil {
		return 0
	}
	if n.Left == nil || n.Right == nil {
		return 1
	}
	s := []*Node{n}
	var level int
	for len(s) > 0 {
		level++
		levelSize := len(s)
		for i := 0; i < levelSize; i++ {
			node := s[0]
			s = s[1:]
			if node.Left == nil && node.Right == nil {
				return level
			}
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
	}
	return level
}

// MaxDepth will return shortest depth of the tree.
func (n *Node) MaxDepth() int {
	if n == nil {
		return 0
	}
	if n.Left == nil || n.Right == nil {
		return 1
	}
	s := []*Node{n}
	var level int
	for len(s) > 0 {
		level++
		levelSize := len(s)
		for i := 0; i < levelSize; i++ {
			node := s[0]
			s = s[1:]
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
	}
	return level
}

// LevelOrderSuccesor will return the node after given node.
func (n *Node) LevelOrderSuccesor(v *Node) *Node {
	if n == nil {
		return nil
	}
	s := []*Node{n}
	for len(s) > 0 {
		node := s[0]
		s = s[1:]
		if node.Right != nil {
			s = append(s, node.Right)
		}
		if node.Left != nil {
			s = append(s, node.Left)
		}
		if node.Value == v.Value {
			break
		}
	}
	return s[0]
}

func (n *Node) traverse() {
	if n != nil {
		s := []*Node{n}
		for len(s) > 0 {
			node := s[0]
			var val int
			if node.Next != nil {
				val = node.Next.Value
			}
			fmt.Printf(strconv.Itoa(node.Value) + " " + strconv.Itoa(val) + " ")
			s = s[1:]
			if node.Right != nil {
				s = append(s, node.Right)
			}
			if node.Left != nil {
				s = append(s, node.Left)
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

func subsets(arr []int) [][]int {
	res := [][]int{{}}
	for _, e := range arr {
		for _, v := range res {
			res = append(res, append(v, e))
		}
	}
	return res
}

func subsetWithDuplicates(arr []int) [][]int {
	res := [][]int{{}}
	var si, ei int
	sort.Ints(arr)
	for idx := range arr {
		si = 0
		if idx > 0 && arr[idx] == arr[idx-1] {
			si = ei +1
		}
		ei = len(res) -1
		for i:=si; i<ei+1; i++ {
			res = append(res, append(res[i], arr[idx]))
		}
	}
	return res
}

func permutations(arr []int) [][]interface{} {
	nLen := len(arr)
	var res [][]interface{}
	var perms [][]interface{}
	perms = append(perms, []interface{}{})
	for _, e := range arr {
		n := len(perms)
		for i:=0; i<n; i++ {
			old := utils.Remove(perms, 0)
			for j:=0; j<len(old)+1; j++ {
				new := utils.Insert(old, j, e)
				if len(new) == nLen {
					res = append(res, new)
				} else {
					perms = append(perms, new)
				}
			}
		}
	}
	return res
}

func permutationStringCase(s string) []string {
	var res []string
	res = append(res, s)
	for i, r := range s {
		if unicode.IsLetter(r) {
			n := len(res)
			for j:=0; j<n; j++ {
				chs := strings.Split(res[j], "")
				s := utils.SwapCase([]rune(chs[i]))
				swip := fmt.Sprintf("%c", s)
				chs[i] = swip
				res = append(res, strings.Join(chs, ""))
			}
		}
	}
	return res
}

func main() {
	// node := &Node{Value: 3}
	root := &Node{Value: 1}
	root.Left = &Node{Value: 3}
	root.Right = &Node{Value: 2}
	root.Left.Left = &Node{Value: 7}
	root.Right.Left = &Node{Value: 5}
	root.Right.Right = &Node{Value: 4}
	root.Left.Right = &Node{Value: 6}
	// root.Right.Left.Right = &Node{Value: 17}
	// fmt.Println(root.TraversePrint())
	// fmt.Println(root.TraversePrintReverse())
	// fmt.Println(root.TraverseZigZagPrint())
	// fmt.Println(root.TraverseAvgPrint())
	// fmt.Println(root.MinDepth())
	// fmt.Println(root.MaxDepth())
	// fmt.Println(root.LevelOrderSuccesor(node))
	// root.ConnectLevelOrder()
	// root.ConnectAllLevelOrder()
	// root.traverse()
	// fmt.Println(root.RightView())
	// fmt.Println(root.TraversePrint())
	// root.traverse()
	// root.traverse()
	// root.postorder()
	// fmt.Println(subsets([]int{1, 5, 3}))
	// fmt.Println(subsetWithDuplicates([]int{1, 5, 3, 3}))
	// fmt.Println(permutations([]int{5,3,1}))
	fmt.Println(permutationStringCase("ad52"))
}
