package main

import (
	"fmt"

	"github.com/ef-ds/deque"
)

func findSubarray(a []int, t int) []deque.Deque {
	var res []deque.Deque
	var res2 [][]int
	p := 1
	l := 0
	for r := range a {
		p*=a[r]
		for p >= t && l < len(a) {
			p/=a[l]
			l++
		}

		var d deque.Deque
		var t []int
		for i := r; i > l-1; i-- {
			d.PushBack(a[i])
			t = append([]int{a[i]}, t...)
			res2 = append(res2, t)
			res = append(res, d)
		}
	}
	fmt.Println(res2)
	return res
}
