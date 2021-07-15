package main

import (
	"github.com/ahrav/learngo/algos/utils"
)

func lengthLongestSubArray(a []int, k int) int {
	m := make(map[int]int)
	ws, l, cc := 0, 0, 0
	for we := range a {
		rc := a[we]
		m[rc]++

		cc = utils.Max(cc, m[rc])

		if we - ws + 1 - cc > k {
			// shrink window
			lc := a[ws]
			m[lc]--
			ws++
		}

		l = utils.Max(we - ws + 1, l)
	}
	return l
}
