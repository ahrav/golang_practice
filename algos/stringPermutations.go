package main

import (
	"strings"

	"github.com/ahrav/learngo/algos/utils"
)

func stringPermutations(s, p string) bool {
	ws, l, pl := 0, 0, len(p)
	for we := range s {
		rc := string(s[we])
		if !strings.Contains(p, rc) {
			ws = we + 1
		}
		l = utils.Max(l, we - ws + 1)
		if l >= pl {
			return true
		}
	}
	return false
}

func stringPermutations2(s, p string) bool {
	ws, ma, l := 0, 0, len(p)
	m := make(map[string]int)
	for c := range p {
		m[string(p[c])]++
	}

	for we := range s {
		rc := string(s[we])
		if _, ok := m[rc]; ok {
			m[rc]--
			if m[rc] == 0 {
				ma++
			}
		}

		if ma == l {
			return true
		}

		if we >= l -1 {
			lc := string(s[ws])
			ws++
			if val, ok := m[lc]; ok {
				if val == 0 {
					ma--
				}
				m[lc]++
			}
		}
	}
	return false
}
