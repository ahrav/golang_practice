package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func main() {
	// myList := linkedList{}
	middleList := linkedList{}
	// node1 := &node{data: 6}
	node2 := &node{data: 5}
	node3 := &node{data: 4}
	node4 := &node{data: 3}
	node5 := &node{data: 2}
	node6 := &node{data: 1}
	// middleList.prepend(node1)
	middleList.prepend(node2)
	middleList.prepend(node3)
	middleList.prepend(node4)
	middleList.prepend(node5)
	middleList.prepend(node6)
	fmt.Println(*middleList.middle())
	// node1.next = myList.head.next
	// myList.printListData()
	// fmt.Println(myList.hasCycle())
	// fmt.Println(myList.cycleLength())
	// fmt.Println(myList.cycleStart())
	newList := linkedList{}
	n1 := &node{data: 12}
	newList.prepend(n1)
	fmt.Println(newList.isHappyNum())
	fmt.Println(findHappyNum(12))
}
