package main

func shortestSubstring(s, p string) string {
	var res []int
	ma, ws, pl, ml, ss := 0, 0, len(p), len(s) + 1, 0
	m := make(map[string]int)

	for c := range p {
		m[string(p[c])]++
	}

	for we := range s {
		rc := string(s[we])
		res = append(res, we)
		if _, ok := m[rc]; ok {
			m[rc]--
			if m[rc] >= 0 {
				ma++
			}
		}

		for ma == pl {
			if ml > we - ws + 1 {
				ml = we - ws + 1
				ss = ws
			}
			lc := string(s[ws])
			res = res[1:]
			ws++
			if val, ok := m[lc]; ok {
				if val == 0 {
					ma--
				}
				m[lc]++
			}
		}
	}
	if ml > len(s) {
		return ""
	}
	return s[ss:ss + ml]
}
