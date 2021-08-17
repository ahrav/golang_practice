package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) prepend (n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) delete(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	prev := l.head
	for prev.next.data != value {
		if prev.next.next == nil {
			return
		}
		prev = prev.next
	}
	prev.next = prev.next.next
	l.length--
}

func (l *linkedList) hasCycle() bool {
	sp, fp := l.head, l.head
	for fp != nil && fp.next != nil {
		fp = fp.next.next
		sp = sp.next
		if fp == sp {
			return true
		}
	}
	return false
}

func (l *linkedList) cycleLength() int {
	sp, fp := l.head, l.head
	for fp != nil && fp.next != nil {
		fp = fp.next.next
		sp = sp.next
		if fp == sp {
			return calcCycleLen(sp)
		}
	}
	return 0
}

func calcCycleLen(n *node) int {
	curr := n
	cnt := 0
	for {
		curr = curr.next
		cnt++
		if curr == n {
			break
		}
	}
	return cnt
}

func (l *linkedList) cycleStartSlow() *node {
	m := make(map[*node]int)
	curr := l.head
	for {
		m[curr]++
		if val, ok := m[curr]; ok {
			if val == 2 {
				return curr
			}
		}
	}
}

func (l *linkedList) cycleStart() *node {
	len := l.cycleLength()
	sp, fp := l.head, l.head
	for i:=0; i<len; i++ {
		fp = fp.next
	}
	for fp != sp {
		sp = sp.next
		fp = fp.next
	}
	return fp
}

func (l *linkedList) isHappyNum() bool {
	m := make(map[node]int)
	curr := l.head
	m[*curr]++
	for {
		str := strconv.Itoa(curr.data)
		arr := strings.Split(str, "")
		sum := 0
		for _, val := range arr {
			i, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Inalid character")
			}
			sum+=int(math.Pow(float64(i), 2))
		}
		if sum == 1 {
			return true
		}
		node := &node{data: sum}
		if _, ok := m[*node]; ok {
			return false
		} else {
			m[*node]++
			curr.next = node
			curr = curr.next
		}
	}
}

func (l *linkedList) middle() *node{
	sp, fp := l.head, l.head
	if fp == nil || fp.next == nil {
		return fp
	}
	for fp != nil && fp.next != nil {
		fp = fp.next.next
		sp = sp.next
	}
	return sp
}

func (l *linkedList) isPalindrome() bool {
	if l.head == nil && l.head.next == nil {
		return true
	}

	middle := l.middle()
	secondHead := reverse(middle)
	copySecondHead := secondHead

	for l.head != nil && secondHead != nil {
		if l.head.data != secondHead.data {
			break
		}
		l.head = l.head.next
		secondHead = secondHead.next
	}
	reverse(copySecondHead)

	if l.head == nil || secondHead == nil {
		return true
	}
	return false
}

func (l *linkedList) reorder() {
	if l.head == nil && l.head.next == nil {
		return
	}
	middle := l.middle()
	secondHead := reverse(middle)
	firstHead := l.head
	for firstHead != nil && secondHead != nil {
		fnxt := firstHead.next
		firstHead.next = secondHead
		firstHead = fnxt
		lnxt := secondHead.next
		secondHead.next = firstHead
		secondHead = lnxt
	}
	if firstHead != nil {
		firstHead.next = nil
	}
}

func reverse(n *node) *node {
	var prev *node
	for n != nil {
		nxt := n.next
		n.next = prev
		prev = n
		n = nxt
	}
	return prev
}

func reverseSublist(n *node, p, q int) *node {
	if p == q {
		return n
	}
	var prev *node
	cur := n
	var i int
	for n != nil && i < p-1 {
		prev = cur
		cur = cur.next
		i++
	}
	lastNodeFirst := prev
	lastNodeSublist := cur
	var next *node

	i = 0
	for cur != nil && i < q-p+1 {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
		i++
	}

	if lastNodeFirst != nil {
		lastNodeFirst.next = prev
	} else {
		n = prev
	}

	lastNodeSublist.next = cur
	return n
}

func reverseKSublist(n *node, k int) *node {
	if k <= 1 {
		return n
	}
	var prev *node
	cur := n
	for {
		lastNodeFirst := prev
		lastNodeSublist := cur
		var next *node
		var i int
		for cur != nil && i < k {
			next = cur.next
			cur.next = prev
			prev = cur
			cur = next
			i++
		}
		if lastNodeFirst != nil {
			lastNodeFirst.next = prev
		} else {
			n = prev
		}
		lastNodeSublist.next = cur
		if cur == nil {
			break
		}
		prev = lastNodeSublist
	}
	return n
}

func reverseKAltSublist(n *node, k int) *node {
	if k <= 1 || n == nil {
		return n
	}
	var prev *node
	cur := n
	for cur != nil {
		lastNodeFirst := prev
		lastNodeSublist := cur
		var next *node
		var i int
		for cur != nil && i < k {
			next = cur.next
			cur.next = prev
			prev = cur
			cur = next
			i++
		}
		if lastNodeFirst != nil {
			lastNodeFirst.next = prev
		} else {
			n = prev
		}
		lastNodeSublist.next = cur
		i = 0
		for cur != nil && i < k {
			prev = cur
			cur = cur.next
			i++
		}
	}
	return n
}

func rotate(n *node, k int) *node {
	start := time.Now()
	if n == nil || n.next == nil || k == 0 {
		return n
	}
	l := 1
	last := n
	for last.next != nil {
		last = last.next
		l++
	}
	k %= l
	var prev *node
	cur := n
	var i int
	for cur != nil && i < k {
		prev = cur
		cur = cur.next
		i++
	}
	lastNodeFirst := prev
	startNodeLast := cur
	for cur != nil {
		prev = cur
		cur = cur.next
	}
	lastNodeFirst.next = cur
	prev.next = n
	n = startNodeLast
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time %s", elapsed)
	return n
}

func rotate2(n *node, k int) *node {
	start := time.Now()
	if n == nil || n.next == nil || k == 0 {
		return n
	}
	l := 1
	last := n
	for last.next != nil {
		last = last.next
		l++
	}
	last.next = n
	k %= l
	skip := l-k
	lastNodeRotated := n
	for i:=0; i<(skip-1); i++ {
		lastNodeRotated = lastNodeRotated.next
	}
	n = lastNodeRotated.next
	lastNodeRotated.next = nil
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time %s", elapsed)
	return n
}

func findHappyNum(n int) bool {
	sp, fp := n, n
	for {
		sp = squareSum(sp)
		fp = squareSum(squareSum(fp))
		if sp == fp {
			break
		}
	}
	return sp == 1
}

func squareSum(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum+=int(math.Pow(float64(d), 2))
		n/=10
	}
	return sum
}

func circularArrLoopExists(arr []int) bool {
	for i := range arr {
		isForward := arr[i] >= 0
		sp, fp := i, i
		for {
			sp = nextIdx(arr, isForward, sp)
			fp = nextIdx(arr, isForward, fp)

			if fp != -1 {
				fp = nextIdx(arr, isForward, fp)
			}
			if sp == -1 || fp == -1 || fp == sp {
				break
			}
		}
		if sp != -1 && sp == fp {
			return true
		}
	}
	return false
}

func nextIdx(arr []int, isForward bool, i int) int {
	dir := arr[i] >= 0

	if isForward != dir {
		return -1
	}

	nextIdx := (i + arr[i]) % len(arr)

	if nextIdx == i {
		nextIdx = -1
	}

	return nextIdx
}

func main() {
	// myList := linkedList{}
	middleList := linkedList{}
	// node1 := &node{data: 7}
	// node2 := &node{data: 6}
	node3 := &node{data: 5}
	node4 := &node{data: 4}
	node5 := &node{data: 3}
	node6 := &node{data: 2}
	node7 := &node{data: 1}
	// middleList.prepend(node1)
	// middleList.prepend(node2)
	middleList.prepend(node3)
	middleList.prepend(node4)
	middleList.prepend(node5)
	middleList.prepend(node6)
	middleList.prepend(node7)
	// middleList.reorder()
	middleList.printListData()
	// x := reverseKSublist(node7, 2)
	x := rotate(node7, 8)
	fmt.Println(x.data, x.next.data, x.next.next.data, x.next.next.next.data, x.next.next.next.next.data)
	// x2 := rotate2(node7, 8)
	// fmt.Println(x2.data, x2.next.data, x2.next.next.data, x2.next.next.next.data, x2.next.next.next.next.data)
	// middleList.printListData()
	// fmt.Println(*middleList.middle())
	// node1.next = myList.head.next
	// myList.printListData()
	// fmt.Println(myList.hasCycle())
	// fmt.Println(myList.cycleLength())
	// fmt.Println(myList.cycleStart())
	// newList := linkedList{}
	// n1 := &node{data: 12}
	// newList.prepend(n1)
	// fmt.Println(newList.isHappyNum())
	// fmt.Println(findHappyNum(12))
	// fmt.Println(circularArrLoopExists([]int{2, 1, -1, -2}))
}
