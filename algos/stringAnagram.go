package main

func stringAnagram(s, p string) []int {
	wl, ws, pl := 0, 0, len(p)
	m := make(map[string]int)
	var res []int

	for we := range string(s) {
		rc := string(s[we])
		m[rc]++
		wl++
		res = append(res, we)
		if len(m) == pl {
			return res
		}
		for we - ws + 1 >= pl {
			lc := string(p[ws])
			ws++
			wl--
			m[lc]--
			res = res[1:]
			if val, ok := m[lc]; ok {
				if val == 0 {
					delete(m, lc)
				}
			}
		}
	}
	return res
}
