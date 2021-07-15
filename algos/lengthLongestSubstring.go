package main

import (
	"github.com/ahrav/learngo/algos/utils"
)

func lengthLongestSubstring(s string, k int) int {
	m := make(map[byte]int)
	ws, l, cc := 0, 0, 0
	for we := range s {
		rc := s[we]
		m[rc]++

		cc = utils.Max(cc, m[rc])

		if we - ws + 1 - cc > k {
			// shrink window
			lc := s[ws]
			m[lc]--
			ws++
		}

		l = utils.Max(we - ws + 1, l)
	}
	return l
}
