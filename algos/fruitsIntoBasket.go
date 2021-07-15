package main

import "github.com/ahrav/learngo/algos/utils"

func fruitsIntoBasket(f []string) int {
	m := make(map[string]int)
	wst := 0
	t := 0
	for we := range f {
		rc := f[we]
		if _, ok := m[rc]; ok {
			m[rc]++
		} else {
			m[rc] = 0
		}

		for len(m) > 2 {
			lc := f[wst]
			if val, ok := m[lc]; ok {
				if val == 0 {
					delete(m, lc)
				} else {
					m[lc]--
				}
			}
			wst++
		}
		t = utils.Max(t, we - wst + 1)
	}
	return t
}
