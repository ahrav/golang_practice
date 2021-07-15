package main

import (
	"github.com/ahrav/learngo/algos/utils"
)

func nonRepeatSubstring(s string) int {
	m := make(map[byte]int)
	wst := 0
	l := 0
	for we := range s {
		rc := s[we]
		if _, ok := m[rc]; ok {
			for i := wst; i < we; i++ {
				lc := s[i]
				if val, ok := m[lc]; ok {
					if val == 0 {
						delete(m, lc)
					}
				}
				wst++
			}
		} else {
			m[rc] = 1
		}
		l = utils.Max(l, we - wst + 1)
	}
	return l
}
