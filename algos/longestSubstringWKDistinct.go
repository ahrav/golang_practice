package main

import "github.com/ahrav/learngo/algos/utils"

func longestSubstringWKDistinct(s string, k int) int {
	m := make(map[byte]int)
	wst := 0
	l := 0
	for we := range s {
		rc := s[we]
		if _, ok := m[rc]; ok {
			m[rc]++
		} else {
			m[rc] = 0
		}
		for len(m) > k {
			lc := s[wst]
			if val, ok := m[lc]; ok {
				if val == 0 {
					delete(m, lc)
				} else {
					m[lc]--
				}
			}
			wst++
		}
		l = utils.Max(l, we - wst + 1)
	}
	return l
}
